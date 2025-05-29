<template>
<div class="max-w-full overflow-x-auto custom-scrollbar">
<div id="chartOne" class="-ml-5 min-w-[650px] xl:min-w-full pl-2">
    <div ref="chart" class="w-[1000px] h-full flex justify-center items-center"></div>

</div>
</div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import VueApexCharts from 'vue3-apexcharts'
import * as d3 from 'd3'

const chart = ref(null)

    onMounted(async () => {
            // 1. 请求数据
            const res = await fetch('http://zzh.hengtai119.cn/api/ageinfo')
            const data = await res.json()


            // 2. 提取 x 和 y
            const margin = { top: 50, right: 50, bottom: 50, left: 50 }

            const width = 600 - margin.left - margin.right
            const height = 300 - margin.top - margin.bottom

            const svg = d3.select(chart.value)
            .append('svg')
            .attr('width', width + margin.left + margin.right)
            .attr('height', height + margin.top + margin.bottom)
            .append('g')
            .attr('transform', `translate(${margin.left},${margin.top})`)

            const x = d3.scaleLinear()
            .domain(d3.extent(data, d => d.age))
            .range([0, width])

            const y = d3.scaleLinear()
            .domain([0, d3.max(data, d => d.count)])
            .range([height, 0])

            // X 轴
            svg.append('g')
            .attr('transform', `translate(0,${height})`)
            .call(d3.axisBottom(x).tickFormat(d3.format('d')))

            // Y 轴
            svg.append('g')
            .call(d3.axisLeft(y))

            // 折线图
            const line = d3.line()
            .x(d => x(d.age))
            .y(d => y(d.count))

            svg.append('path')
            .datum(data)
            .attr('fill', 'none')
            .attr('stroke', 'steelblue')
            .attr('stroke-width', 2)
            .attr('d', line)

            // 添加 tooltip div
            const tooltip = d3.select(chart.value)
            .append('div')
            .style('position', 'absolute')
            .style('background', 'rgba(0,0,0,0.7)')
            .style('color', '#fff')
            .style('padding', '5px 8px')
            .style('border-radius', '4px')
            .style('font-size', '12px')
            .style('pointer-events', 'none')
            .style('display', 'none')

            // 加圆点并绑定事件
            svg.selectAll('circle')
            .data(data)
            .enter()
            .append('circle')
            .attr('cx', d => x(d.age))
            .attr('cy', d => y(d.count))
            .attr('r', 4)
            .attr('fill', 'orange')
            .on('mouseover', (event, d) => {
                    tooltip.style('display', 'block')
                    .html(`年龄: <strong>${d.age}</strong><br>数量: <strong>${d.count}</strong>`)
                    })
            .on('mousemove', event => {
                    tooltip.style('left', (event.pageX + 10) + 'px')
                    .style('top', (event.pageY + 10) + 'px')
                    })
            .on('mouseout', () => {
                    tooltip.style('display', 'none')
                    })
    })
</script>
