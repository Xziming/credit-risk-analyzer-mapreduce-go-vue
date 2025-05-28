<template>
<div ref="chart" class="w-full h-[400px] flex justify-center items-center"></div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import * as d3 from 'd3'

const chart = ref(null)

    onMounted(async () => {
            const res = await fetch('http://zzh.hengtai119.cn/api/limitbalinfo')
            const raw = await res.json()

            // 将 range 转为数字并排序
            const data = raw.map(d => ({
range: +d.range, // 转为数字
count: d.count
})).sort((a, b) => a.range - b.range)

            const margin = { top: 30, right: 30, bottom: 60, left: 60 }
            const width = 650 - margin.left - margin.right
            const height = 400 - margin.top - margin.bottom

            const svg = d3.select(chart.value)
            .append('svg')
            .attr('width', width + margin.left + margin.right)
            .attr('height', height + margin.top + margin.bottom)
            .append('g')
            .attr('transform', `translate(${margin.left},${margin.top})`)

    // X 轴：额度范围
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

    // Y 轴：人数
    const y = d3.scaleLinear()
    .domain([0, d3.max(data, d => d.count)])
    .nice()
.range([height, 0])

    svg.append('g').call(d3.axisLeft(y))

    // 柱状图动画
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
    .attr('y', d => y(d.count))
    .attr('height', d => height - y(d.count))

    // 标签文字
    svg.selectAll('text.label')
    .data(data)
.enter()
    .append('text')
    .attr('class', 'label')
    .attr('x', d => x(d.range) + x.bandwidth() / 2)
    .attr('y', d => y(d.count) - 5)
.text(d => d.count)
    .style('text-anchor', 'middle')
    .style('font-size', '12px')
    .style('fill', '#333')
    })
</script>

<style scoped>
svg {
    font-family: sans-serif;
}
</style>

