<template>
  <div ref="chart" class="w-full h-[500px]"></div>
</template>

<script setup>
import * as d3 from 'd3'
import { ref, onMounted } from 'vue'
import axios from 'axios'

const chart = ref()

onMounted(async () => {
  // 1. 获取数据
  const response = await axios.get('http://zzh.hengtai119.cn/api/limitdefaults')
  const rawData = response.data

  // 2. 预处理数据
  const grouped = {}
  rawData.forEach(item => {
    if (!grouped[item.range]) {
      grouped[item.range] = { range: item.range, defaulted: 0, nonDefaulted: 0 }
    }
    if (item.defaulted === 1) {
      grouped[item.range].defaulted = item.count
    } else {
      grouped[item.range].nonDefaulted = item.count
    }
  })

  const data = Object.values(grouped)

  // 3. 绘图参数
  const margin = { top: 40, right: 30, bottom: 50, left: 60 }
  const width = 800 - margin.left - margin.right
  const height = 500 - margin.top - margin.bottom

  const svg = d3.select(chart.value)
    .append('svg')
    .attr('width', width + margin.left + margin.right)
    .attr('height', height + margin.top + margin.bottom)
    .append('g')
    .attr('transform', `translate(${margin.left},${margin.top})`)

  const subgroups = ['defaulted', 'nonDefaulted']
  const groups = data.map(d => d.range)

  const x = d3.scaleBand()
    .domain(groups)
    .range([0, width])
    .padding(0.2)

  svg.append('g')
    .attr('transform', `translate(0, ${height})`)
    .call(d3.axisBottom(x))
    .selectAll('text')
    .attr('transform', 'rotate(-45)')
    .style('text-anchor', 'end')

  const y = d3.scaleLinear()
    .domain([0, d3.max(data, d => d.defaulted + d.nonDefaulted)])
    .nice()
    .range([height, 0])

  svg.append('g')
    .call(d3.axisLeft(y))

  const color = d3.scaleOrdinal()
    .domain(subgroups)
    .range(['#e41a1c', '#377eb8'])

  const stackedData = d3.stack().keys(subgroups)(data)

  svg.append('g')
    .selectAll('g')
    .data(stackedData)
    .join('g')
    .attr('fill', d => color(d.key))
    .selectAll('rect')
    .data(d => d)
    .join('rect')
    .attr('x', d => x(d.data.range))
    .attr('y', d => y(d[1]))
    .attr('height', d => y(d[0]) - y(d[1]))
    .attr('width', x.bandwidth())

  svg.append('text')
    .attr('x', width / 2)
    .attr('y', -10)
    .attr('text-anchor', 'middle')
    .attr('font-size', '16px')
    .text('信用额度 vs 违约与非违约人数')
})
</script>

