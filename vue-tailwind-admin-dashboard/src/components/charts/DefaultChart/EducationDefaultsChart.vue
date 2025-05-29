<template>
<div class="max-w-full overflow-x-auto custom-scrollbar">
<div id="educationDefaultsChart" class="-ml-5 min-w-[650px] xl:min-w-full pl-2">
<div ref="chart" class="w-[1000px] h-full flex justify-center items-center"></div>

</div>
</div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import VueApexCharts from 'vue3-apexcharts'
import * as d3 from 'd3'

const chart = ref(null)

    async function fetchData() {
        // 替换成你实际的接口地址
        const res = await fetch('http://zzh.hengtai119.cn/api/educationdefaults')
            if (!res.ok) throw new Error('网络请求失败')
                return await res.json()
    }

onMounted(async () => {
        let raw = []
        try {
        raw = await fetchData()
        } catch (error) {
        console.error('获取数据失败:', error)
        return
        }

        const margin = { top: 30, right: 30, bottom: 80, left: 60 }
        const width = 700 - margin.left - margin.right
        const height = 400 - margin.top - margin.bottom

        const educationMap = {
1: '高中及以下',
2: '大学专科',
3: '大学本科',
4: '硕士及以上'
}

// 整理数据格式
const grouped = new Map()
    raw.forEach(d => {
            const edu = educationMap[d.education] || '未知'
            if (!grouped.has(edu)) grouped.set(edu, { education: edu, defaulted0: 0, defaulted1: 0 })
            const obj = grouped.get(edu)
            if (d.defaulted === 0) obj.defaulted0 = d.count
            else obj.defaulted1 = d.count
            })

const data = Array.from(grouped.values())

    const subgroups = ['defaulted0', 'defaulted1']

    const color = d3.scaleOrdinal()
.domain(subgroups)
    .range(['#42b883', '#ff6b6b'])

const svg = d3.select(chart.value)
    .append('svg')
    .attr('width', width + margin.left + margin.right)
    .attr('height', height + margin.top + margin.bottom)
    .append('g')
    .attr('transform', `translate(${margin.left},${margin.top})`)

    const x = d3.scaleBand()
    .domain(data.map(d => d.education))
    .range([0, width])
.padding([0.2])

    svg.append('g')
    .attr('transform', `translate(0,${height})`)
.call(d3.axisBottom(x))
    .selectAll("text")
    .attr("transform", "rotate(-40)")
    .style("text-anchor", "end")

    const xSubgroup = d3.scaleBand()
    .domain(subgroups)
    .range([0, x.bandwidth()])
.padding([0.05])

    const y = d3.scaleLinear()
    .domain([0, d3.max(data, d => Math.max(d.defaulted0, d.defaulted1)) * 1.1])
    .nice()
.range([height, 0])

    svg.append('g').call(d3.axisLeft(y))

    const groups = svg.selectAll('g.layer')
    .data(data)
.enter()
    .append('g')
    .attr('transform', d => `translate(${x(d.education)},0)`)

    groups.selectAll('rect')
.data(d => subgroups.map(key => ({ key, value: d[key] })))
.enter()
    .append('rect')
    .attr('x', d => xSubgroup(d.key))
    .attr('y', height)
    .attr('width', xSubgroup.bandwidth())
    .attr('height', 0)
    .attr('fill', d => color(d.key))
    .transition()
.duration(1000)
    .attr('y', d => y(d.value))
    .attr('height', d => height - y(d.value))

    groups.selectAll('text.label')
.data(d => subgroups.map(key => ({ key, value: d[key] })))
.enter()
    .append('text')
    .attr('class', 'label')
    .attr('x', d => xSubgroup(d.key) + xSubgroup.bandwidth() / 2)
    .attr('y', d => y(d.value) - 5)
.text(d => d.value)
    .style('text-anchor', 'middle')
    .style('font-size', '12px')
    .style('fill', '#333')

    const legend = svg.append('g')
    .attr('transform', `translate(${width - 140}, 0)`)

    legend.selectAll('rect')
    .data(subgroups)
.enter()
    .append('rect')
    .attr('x', 0)
    .attr('y', (d, i) => i * 25)
    .attr('width', 18)
    .attr('height', 18)
    .attr('fill', d => color(d))

    legend.selectAll('text')
    .data(subgroups)
.enter()
    .append('text')
    .attr('x', 24)
    .attr('y', (d, i) => i * 25 + 13)
    .text(d => d === 'defaulted0' ? '未违约' : '违约')
    .style('font-size', '14px')
    .attr('fill', '#555')
    })
</script>

<style scoped>
svg {
    font-family: sans-serif;
}
</style>

