<template>
        <div class="max-w-full overflow-x-auto custom-scrollbar">
                <div id="lineChart" class="-ml-5 min-w-[650px] xl:min-w-full pl-2">
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



