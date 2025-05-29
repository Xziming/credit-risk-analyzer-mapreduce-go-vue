<template>
        <div class="max-w-full overflow-x-auto custom-scrollbar">
                <div id="marriageDefaultsChart" class="-ml-5 min-w-[650px] xl:min-w-full pl-2">
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
        const res = await fetch('http://zzh.hengtai119.cn/api/marriagedefaults')
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

        const margin = { top: 30, right: 30, bottom: 60, left: 60 }
        const width = 700 - margin.left - margin.right
        const height = 400 - margin.top - margin.bottom

        // 婚姻状态映射（可根据实际调整）
        const marriageMap = {
0: '未知',
1: '未婚',
2: '已婚',
3: '离异'
}

// 整理数据格式，转换成 [{ marriage: '已婚', defaulted0: count, defaulted1: count }, ...]
const grouped = new Map()
    raw.forEach(d => {
            const m = marriageMap[d.marriage] || '其他'
            if (!grouped.has(m)) grouped.set(m, { marriage: m, defaulted0: 0, defaulted1: 0 })
            const obj = grouped.get(m)
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

    // x轴 - 婚姻状态
    const x = d3.scaleBand()
    .domain(data.map(d => d.marriage))
    .range([0, width])
.padding([0.2])

    svg.append('g')
    .attr('transform', `translate(0,${height})`)
.call(d3.axisBottom(x))

    // 组内子类别位置
    const xSubgroup = d3.scaleBand()
    .domain(subgroups)
    .range([0, x.bandwidth()])
.padding([0.05])

    // y轴 - 数量
    const y = d3.scaleLinear()
    .domain([0, d3.max(data, d => Math.max(d.defaulted0, d.defaulted1)) * 1.1])
    .nice()
.range([height, 0])

    svg.append('g').call(d3.axisLeft(y))

    // 绘制柱状条
    const groups = svg.selectAll('g.layer')
    .data(data)
.enter()
    .append('g')
    .attr('transform', d => `translate(${x(d.marriage)},0)`)

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

    // 添加数值标签
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

    // 图例
    const legend = svg.append('g')
    .attr('transform', `translate(${width - 120}, 0)`)

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

