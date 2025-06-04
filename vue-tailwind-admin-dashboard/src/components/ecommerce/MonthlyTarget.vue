<template>
  <div class="bg-white rounded-xl p-6 shadow-sm dark:bg-gray-900 relative">
    <h2 class="text-2xl font-semibold text-gray-800 dark:text-white mb-6">风险等级占比</h2>
    <div v-if="loading" class="w-full h-[400px] flex items-center justify-center">
      <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary"></div>
    </div>
    <div v-else-if="error" class="w-full h-[400px] flex items-center justify-center text-red-500">
      数据加载失败: {{ error }}
    </div>
    <div v-else ref="svgRef" class="w-full h-[400px]" />
    <div ref="tooltip" class="tooltip hidden absolute px-2 py-1 rounded bg-gray-800 text-white text-sm pointer-events-none"></div>
  </div>
</template>

<script setup>
import * as d3 from 'd3'
import { onMounted, ref } from 'vue'

const svgRef = ref(null)
const tooltip = ref(null)
const loading = ref(true)
const error = ref(null)
const groupedData = ref([])

const fetchData = async () => {
  try {
    loading.value = true
    error.value = null
    const response = await fetch('http://zzh.hengtai119.cn/api/overdue')
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    const data = await response.json()
    
    // Group data into risk categories
    const noOverdue = data.find(item => item.max_overdue_streak === 0)?.customer_count || 0
    const mildOverdue = data
      .filter(item => item.max_overdue_streak >= 1 && item.max_overdue_streak <= 2)
      .reduce((sum, item) => sum + item.customer_count, 0)
    const highRisk = data
      .filter(item => item.max_overdue_streak >= 3)
      .reduce((sum, item) => sum + item.customer_count, 0)
    
    groupedData.value = [
      { label: '无逾期', value: noOverdue, color: '#10B981' },
      { label: '轻度逾期', value: mildOverdue, color: '#F59E0B' },
      { label: '高风险', value: highRisk, color: '#EF4444' },
    ]
  } catch (err) {
    error.value = err.message
    console.error('Error fetching data:', err)
  } finally {
    loading.value = false
  }
}

const drawChart = () => {
  if (loading.value || error.value || !groupedData.value.length) return
  
  const container = svgRef.value
  if (!container) return
  container.innerHTML = ''

  const width = container.offsetWidth
  const height = 400
  const radius = Math.min(width, height) / 2 - 20

  const svg = d3.select(container)
    .append('svg')
    .attr('width', width)
    .attr('height', height)
    .append('g')
    .attr('transform', `translate(${width / 2}, ${height / 2})`)

  const pie = d3.pie()
    .sort(null)
    .value(d => d.value)

  const arc = d3.arc()
    .innerRadius(radius * 0.6)  // 设为60%半径，变成圆环
    .outerRadius(radius)

  const arcs = svg.selectAll('arc')
    .data(pie(groupedData.value))
    .enter()
    .append('g')

  arcs.append('path')
    .attr('d', arc)
    .attr('fill', d => d.data.color)
    .attr('stroke', 'white')
    .attr('stroke-width', 2)
    .on('mouseover', (event, d) => {
      tooltip.value.classList.remove('hidden')
      tooltip.value.textContent = `${d.data.label}: ${d.data.value} (${((d.data.value / d3.sum(groupedData.value, d => d.value)) * 100).toFixed(1)}%)`
    })
    .on('mousemove', (event) => {
      const [x, y] = d3.pointer(event, container)
      const tooltipWidth = tooltip.value.offsetWidth
      const tooltipHeight = tooltip.value.offsetHeight
      let left = x + 10
      let top = y + 10
      if (left + tooltipWidth > width) {
        left = x - tooltipWidth - 10
      }
      if (top + tooltipHeight > height) {
        top = y - tooltipHeight - 10
      }
      tooltip.value.style.left = `${left}px`
      tooltip.value.style.top = `${top}px`
    })
    .on('mouseout', () => {
      tooltip.value.classList.add('hidden')
    })

  arcs.append('text')
    .attr('transform', d => `translate(${arc.centroid(d)})`)
    .attr('text-anchor', 'middle')
    .attr('alignment-baseline', 'middle')
    .attr('fill', 'white')
    .attr('font-size', '12px')
    .text(d => `${((d.data.value / d3.sum(groupedData.value, d => d.value)) * 100).toFixed(1)}%`)

  // 图例
  const legend = svg.append('g')
    .attr('transform', `translate(${-width / 2 + 20}, ${-height / 2 + 20})`)

  groupedData.value.forEach((d, i) => {
    const g = legend.append('g').attr('transform', `translate(0, ${i * 20})`)
    g.append('rect')
      .attr('width', 12)
      .attr('height', 12)
      .attr('fill', d.color)
    g.append('text')
      .attr('x', 18)
      .attr('y', 10)
      .attr('font-size', '12px')
      .attr('fill', '#4B5563')
      .text(d.label)
  })
}

onMounted(async () => {
  await fetchData()
  drawChart()
})
</script>

<style scoped>
.tooltip {
  position: absolute;
  pointer-events: none;
  white-space: nowrap;
  transition: opacity 0.2s ease;
  opacity: 0.9;
  z-index: 10;
}
svg {
  max-width: 100%;
}
</style>
