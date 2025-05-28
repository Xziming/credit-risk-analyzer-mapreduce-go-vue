<template>
  <div ref="chart" class="w-full h-96 flex justify-center items-center relative"></div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import * as d3 from 'd3'

const chart = ref(null)

onMounted(async () => {
  // 1. 获取数据
  const res = await fetch('http://zzh.hengtai119.cn/api/genderinfo')
  const rawData = await res.json()

  // 2. 转换数据格式：gender 映射为 "男" 和 "女"
  const data = rawData.map(d => ({
    label: d.gender === 1 ? '男' : '女',
    value: d.count
  }))

  // 设置尺寸
  const width = 400
  const height = 400
  const radius = Math.min(width, height) / 2

  const svg = d3.select(chart.value)
    .append('svg')
    .attr('width', width)
    .attr('height', height)
    .append('g')
    .attr('transform', `translate(${width / 2}, ${height / 2})`)

  const color = d3.scaleOrdinal()
    .domain(data.map(d => d.label))
    .range(['#4FC3F7', '#FF8A65'])

  const pie = d3.pie()
    .value(d => d.value)

  const arc = d3.arc()
    .innerRadius(0)
    .outerRadius(radius)

  const tooltip = d3.select(chart.value)
    .append('div')
    .style('position', 'absolute')
    .style('background', 'rgba(0,0,0,0.7)')
    .style('color', '#fff')
    .style('padding', '6px 10px')
    .style('border-radius', '4px')
    .style('font-size', '12px')
    .style('pointer-events', 'none')
    .style('display', 'none')

  // 绘制饼图
  svg.selectAll('path')
    .data(pie(data))
    .enter()
    .append('path')
    .attr('d', arc)
    .attr('fill', d => color(d.data.label))
    .on('mouseover', (event, d) => {
      tooltip.style('display', 'block')
        .html(`${d.data.label}: <strong>${d.data.value}</strong>`)
    })
    .on('mousemove', (event) => {
      tooltip
        .style('left', (event.pageX + 10) + 'px')
        .style('top', (event.pageY + 10) + 'px')
    })
    .on('mouseout', () => {
      tooltip.style('display', 'none')
    })

  // 添加文字标签
  svg.selectAll('text')
    .data(pie(data))
    .enter()
    .append('text')
    .text(d => d.data.label)
    .attr('transform', d => `translate(${arc.centroid(d)})`)
    .attr('text-anchor', 'middle')
    .style('fill', '#fff')
    .style('font-size', '14px')
    .style('font-weight', 'bold')
})
</script>

<style scoped>
svg {
  display: block;
  margin: 0 auto;
}
</style>

