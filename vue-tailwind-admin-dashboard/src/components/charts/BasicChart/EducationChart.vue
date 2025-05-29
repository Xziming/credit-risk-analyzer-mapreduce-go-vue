<template>
        <div class="max-w-full overflow-x-auto custom-scrollbar">
                <div id="educationChart" class="-ml-5 min-w-[650px] xl:min-w-full pl-2">
                        <div ref="chart" class="w-[1000px] h-full flex justify-center items-center"></div>

                </div>
        </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import VueApexCharts from 'vue3-apexcharts'
import * as d3 from 'd3'

const chart = ref(null)

// 教育程度映射（可根据实际数据修改）
const educationMap = {
  0: '未知',
  1: '研究生',
  2: '大学',
  3: '高中',
  4: '中学',
  5: '小学',
  6: '其他'
}

onMounted(async () => {
  // 获取数据
  const res = await fetch('http://zzh.hengtai119.cn/api/educationinfo')
  const data = await res.json()

  // 设置边距与画布尺寸
  const margin = { top: 40, right: 40, bottom: 40, left: 120 }
  const width = 650 - margin.left - margin.right
  const height = 400 - margin.top - margin.bottom

  // 创建 SVG 容器
  const svg = d3.select(chart.value)
    .append('svg')
    .attr('width', width + margin.left + margin.right)
    .attr('height', height + margin.top + margin.bottom)
    .append('g')
    .attr('transform', `translate(${margin.left},${margin.top})`)

  // Y 轴（教育等级）
  const y = d3.scaleBand()
    .domain(data.map(d => educationMap[d.education]))
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

  // 绘制条形图，添加动画
  svg.selectAll('rect')
    .data(data)
    .enter()
    .append('rect')
    .attr('y', d => y(educationMap[d.education]))
    .attr('height', y.bandwidth())
    .attr('x', 0)
    .attr('width', 0) // 初始宽度为 0
    .attr('fill', '#4f8ef7')
    .transition()
    .duration(1000)
    .attr('width', d => x(d.count))

  // 添加数量文本标签
  svg.selectAll('text.label')
    .data(data)
    .enter()
    .append('text')
    .attr('class', 'label')
    .attr('x', d => x(d.count) + 5)
    .attr('y', d => y(educationMap[d.education]) + y.bandwidth() / 2 + 4)
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

