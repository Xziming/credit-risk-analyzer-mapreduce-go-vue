<template>
<div class="max-w-full overflow-x-auto custom-scrollbar">
<div id="limitChart" class="-ml-5 min-w-[1500px] xl:min-w-full pl-2">
<div ref="chart" class="w-[1500px] h-full flex justify-center items-center"></div>

</div>
</div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import * as d3 from 'd3'

const chart = ref(null)

onMounted(async () => {
  const res = await fetch('http://zzh.hengtai119.cn/api/limitbalinfo')
  const raw = await res.json()

  // 数据预处理
  const data = raw.map(d => ({
    range: +d.range,
    count: d.count,
  })).sort((a, b) => a.range - b.range)

  const totalCount = d3.sum(data, d => d.count)
  data.forEach(d => {
    d.rate = +(d.count / totalCount).toFixed(4) // 比例数据用于折线图
  })

  const margin = { top: 30, right: 50, bottom: 60, left: 60 }
  const width = 1200 - margin.left - margin.right
  const height = 600 - margin.top - margin.bottom

  const svg = d3.select(chart.value)
    .append('svg')
    .attr('width', width + margin.left + margin.right)
    .attr('height', height + margin.top + margin.bottom)
    .append('g')
    .attr('transform', `translate(${margin.left},${margin.top})`)

  // X 轴
  const x = d3.scaleBand()
    .domain(data.map(d => d.range))
    .range([0, width])
    .padding(0.2)

  svg.append('g')
    .attr('transform', `translate(0,${height})`)
    .call(d3.axisBottom(x).tickFormat(d3.format(',')))
    .selectAll('text')
    .attr('transform', 'rotate(60)')
    .style('text-anchor', 'start')

  // Y 左轴：柱状图（人数）
  const yLeft = d3.scaleLinear()
    .domain([0, d3.max(data, d => d.count)]).nice()
    .range([height, 0])

  svg.append('g').call(d3.axisLeft(yLeft))

  // Y 右轴：折线图（比例）
  const yRight = d3.scaleLinear()
    .domain([0, d3.max(data, d => d.rate)]).nice()
    .range([height, 0])

  svg.append('g')
    .attr('transform', `translate(${width},0)`)
    .call(d3.axisRight(yRight).tickFormat(d3.format(".0%")))

  // 柱状图
  svg.selectAll('rect')
    .data(data)
    .enter()
    .append('rect')
    .attr('x', d => x(d.range))
    .attr('width', x.bandwidth())
    .attr('y', height)
    .attr('height', 0)
    .attr('fill', '#42b883')
    .transition()
    .duration(1000)
    .attr('y', d => yLeft(d.count))
    .attr('height', d => height - yLeft(d.count))

  // 折线图路径
  const line = d3.line()
    .x(d => x(d.range) + x.bandwidth() / 2)
    .y(d => yRight(d.rate))

  svg.append('path')
    .datum(data)
    .attr('fill', 'none')
    .attr('stroke', '#f97316') // 橙色
    .attr('stroke-width', 2.5)
    .attr('d', line)

  // 折线图圆点
  svg.selectAll('circle')
    .data(data)
    .enter()
    .append('circle')
    .attr('cx', d => x(d.range) + x.bandwidth() / 2)
    .attr('cy', d => yRight(d.rate))
    .attr('r', 4)
    .attr('fill', '#f97316')

  // 柱状图标签
  svg.selectAll('text.label')
    .data(data)
    .enter()
    .append('text')
    .attr('class', 'label')
    .attr('x', d => x(d.range) + x.bandwidth() / 2)
    .attr('y', d => yLeft(d.count) - 5)
    .text(d => d.count)
    .style('text-anchor', 'middle')
    .style('font-size', '12px')
    .style('fill', '#333')

  // 折线图标签（百分比）
  svg.selectAll('text.rate-label')
    .data(data)
    .enter()
    .append('text')
    .attr('class', 'rate-label')
    .attr('x', d => x(d.range) + x.bandwidth() / 2)
    .attr('y', d => yRight(d.rate) - 10)
    .text(d => (d.rate * 100).toFixed(1) + '%')
    .style('text-anchor', 'middle')
    .style('font-size', '11px')
    .style('fill', '#f97316')
})
</script>

<style scoped>
svg {
    font-family: sans-serif;
}
</style>
