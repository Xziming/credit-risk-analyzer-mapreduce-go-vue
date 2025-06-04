<template>
  <div class="bg-white rounded-xl p-6 shadow-sm dark:bg-gray-900 lg:col-span-2">
    <div class="flex flex-col md:flex-row items-center justify-between mb-6">
      <h2 class="text-2xl font-semibold text-gray-800 dark:text-white mb-4 md:mb-0">逾期时长分布</h2>
      <div class="flex flex-wrap gap-2">
        <select v-model="chartType" class="bg-gray-100 border border-gray-300 text-gray-700 py-2 px-3 rounded-md dark:bg-gray-800 dark:text-white">
          <option value="line">折线图</option>
          <option value="bar">柱状图</option>
        </select>
        <button @click="setScale('linear')"
                :class="['px-3 py-1 text-sm rounded-md', scale === 'linear' ? 'bg-primary text-white' : 'bg-gray-200 text-gray-700 dark:bg-gray-700 dark:text-white']">
          线性轴
        </button>
        <button @click="setScale('log')"
                :class="['px-3 py-1 text-sm rounded-md', scale === 'log' ? 'bg-primary text-white' : 'bg-gray-200 text-gray-700 dark:bg-gray-700 dark:text-white']">
          对数轴
        </button>
      </div>
    </div>
    <div v-if="loading" class="w-full h-[400px] flex items-center justify-center">
      <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary"></div>
    </div>
    <div v-else-if="error" class="w-full h-[400px] flex items-center justify-center text-red-500">
      数据加载失败: {{ error }}
    </div>
    <div v-else ref="chartContainer" class="relative w-full h-[400px]" />
  </div>
</template>

<script setup>
import * as d3 from 'd3'
import { ref, onMounted, watch, nextTick } from 'vue'

const chartContainer = ref(null)
const chartType = ref('line')
const scale = ref('linear')
const loading = ref(true)
const error = ref(null)
const apiData = ref([])

// 颜色函数
const getColor = (months) => {
  if (months === 0) return '#10B981' // green
  if (months >= 1 && months <= 2) return '#F59E0B' // yellow
  return '#EF4444' // red
}

const fetchData = async () => {
  try {
    loading.value = true
    error.value = null
    const response = await fetch('http://zzh.hengtai119.cn/api/overdue')
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    const data = await response.json()
    // Transform data to match our expected format
    apiData.value = data.map(item => ({
      months: item.max_overdue_streak,
      customers: item.customer_count
    }))
  } catch (err) {
    error.value = err.message
    console.error('Error fetching data:', err)
  } finally {
    loading.value = false
  }
}

const drawChart = () => {
  if (loading.value || error.value || !apiData.value.length) return
  
  const container = chartContainer.value
  if (!container) return
  container.innerHTML = ''

  const margin = { top: 20, right: 30, bottom: 50, left: 50 }
  const width = container.offsetWidth
  const height = 400
  const innerWidth = width - margin.left - margin.right
  const innerHeight = height - margin.top - margin.bottom

  const svg = d3.select(container)
    .append('svg')
    .attr('width', width)
    .attr('height', height)
    .append('g')
    .attr('transform', `translate(${margin.left},${margin.top})`)

  const x = d3.scaleBand()
    .domain(apiData.value.map(d => d.months))
    .range([0, innerWidth])
    .padding(0.2)

  const y = scale.value === 'log'
    ? d3.scaleLog().domain([1, d3.max(apiData.value, d => d.customers)]).range([innerHeight, 0])
    : d3.scaleLinear().domain([0, d3.max(apiData.value, d => d.customers)]).range([innerHeight, 0])

  const xAxis = d3.axisBottom(x).tickFormat(d => `${d}月`)
  const yAxis = d3.axisLeft(y).ticks(5)

  // 添加风险等级背景色区域
  const regions = [
    { range: [0, 0], color: '#10B981', label: '无逾期' },
    { range: [1, 2], color: '#F59E0B', label: '轻度逾期' },
    { range: [3, 6], color: '#EF4444', label: '高风险' }
  ]

  regions.forEach(region => {
    const start = x(region.range[0]) ?? 0
    const end = x(region.range[1]) ?? innerWidth
    const endX = x(region.range[1] + 1) ?? innerWidth
    svg.append('rect')
      .attr('x', start)
      .attr('y', 0)
      .attr('width', endX - start)
      .attr('height', innerHeight)
      .attr('fill', region.color)
      .attr('opacity', 0.05)

    svg.append('text')
      .attr('x', start + (endX - start) / 2)
      .attr('y', -8)
      .attr('text-anchor', 'middle')
      .attr('fill', region.color)
      .attr('font-size', '12px')
      .attr('font-weight', 'bold')
      .text(region.label)
  })

  svg.append('g')
    .attr('transform', `translate(0,${innerHeight})`)
    .call(xAxis)

  svg.append('g')
    .call(yAxis)

  if (chartType.value === 'bar') {
    svg.selectAll('rect.bar')
      .data(apiData.value)
      .enter()
      .append('rect')
      .attr('class', 'bar')
      .attr('x', d => x(d.months))
      .attr('y', d => y(d.customers))
      .attr('width', x.bandwidth())
      .attr('height', d => innerHeight - y(d.customers))
      .attr('fill', d => getColor(d.months))
  } else {
    const line = d3.line()
      .x(d => x(d.months) + x.bandwidth() / 2)
      .y(d => y(d.customers))
      .curve(d3.curveMonotoneX)

    svg.append('path')
      .datum(apiData.value)
      .attr('fill', 'none')
      .attr('stroke', '#1E40AF')
      .attr('stroke-width', 3)
      .attr('d', line)

    svg.selectAll('circle')
      .data(apiData.value)
      .enter()
      .append('circle')
      .attr('cx', d => x(d.months) + x.bandwidth() / 2)
      .attr('cy', d => y(d.customers))
      .attr('r', 5)
      .attr('fill', d => getColor(d.months))
      .attr('stroke', 'white')
      .attr('stroke-width', 1.5)
  }
}

const setScale = type => {
  scale.value = type
}

onMounted(async () => {
  await fetchData()
  drawChart()
})

watch([chartType, scale, apiData], async () => {
  await nextTick()
  drawChart()
})
</script>
