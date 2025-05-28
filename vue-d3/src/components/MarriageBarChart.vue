<template>
  <div ref="chart" class="w-full h-[400px] flex justify-center items-center"></div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import * as d3 from 'd3'

const chart = ref(null)

const marriageMap = {
  0: '未知',
  1: '已婚',
  2: '单身',
  3: '其他'
}

onMounted(async () => {
  // 获取数据
  const res = await fetch('http://zzh.hengtai119.cn/api/marriageinfo')
  const data = await res.json()

  // 设置尺寸和边距
  const margin = { top: 40, right: 40, bottom: 40, left: 100 }
  const width = 600 - margin.left - margin.right
  const height = 300 - margin.top - margin.bottom

  // 创建 SVG
  const svg = d3.select(chart.value)
    .append('svg')
    .attr('width', width + margin.left + margin.right)
    .attr('height', height + margin.top + margin.bottom)
    .append('g')
    .attr('transform', `translate(${margin.left},${margin.top})`)

  // Y 轴（婚姻状态）
  const y = d3.scaleBand()
    .domain(data.map(d => marriageMap[d.marriage]))
    .range([0, height])
    .padding(0.2)

  svg.append('g')
    .call(d3.axisLeft(y))

  // X 轴（数量）
  const x = d3.scaleLinear()
    .domain([0, d3.max(data, d => d.count)])
    .range([0, width])

  svg.append('g')
    .attr('transform', `translate(0,${height})`)
    .call(d3.axisBottom(x).ticks(5))

  // 添加条形图 + 动画
  svg.selectAll('rect')
    .data(data)
    .enter()
    .append('rect')
    .attr('y', d => y(marriageMap[d.marriage]))
    .attr('height', y.bandwidth())
    .attr('x', 0)
    .attr('width', 0) // 初始宽度为 0
    .attr('fill', '#69b3a2')
    .transition()
    .duration(800)
    .attr('width', d => x(d.count))

  // 添加文字标签
  svg.selectAll('text.label')
    .data(data)
    .enter()
    .append('text')
    .attr('class', 'label')
    .attr('x', d => x(d.count) + 5)
    .attr('y', d => y(marriageMap[d.marriage]) + y.bandwidth() / 2 + 4)
    .text(d => d.count)
    .style('font-size', '12px')
    .style('fill', '#333')
})
</script>

<style scoped>
svg {
  font-family: sans-serif;
}
</style>

