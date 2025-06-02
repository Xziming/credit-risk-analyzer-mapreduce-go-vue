<template>
  <div class="max-w-full overflow-x-auto custom-scrollbar">
    <div v-if="loading" class="flex justify-center items-center h-64">
      <div class="text-lg text-gray-600">加载中...</div>
    </div>
    <div v-else-if="error" class="flex justify-center items-center h-64">
      <div class="text-lg text-red-600">{{ error }}</div>
    </div>
    <div v-else id="chartOne" class="min-w-[650px] xl:min-w-full">
      <div ref="chart" class="w-full h-[500px] flex justify-center items-center"></div>
      <div class="mt-4 flex items-center space-x-4">
        <label for="month-filter" class="font-semibold">选择月份:</label>
        <select id="month-filter" v-model="selectedMonth" @change="handleMonthChange" class="border rounded px-2 py-1">
          <option value="all">全部</option>
          <option v-for="month in months" :key="month" :value="month">{{ month }}</option>
        </select>
      </div>
      <table class="mt-4 w-full table-auto border-collapse border border-gray-200 text-sm">
        <thead>
          <tr class="bg-gray-100">
            <th class="border border-gray-300 px-4 py-2">月份</th>
            <th class="border border-gray-300 px-4 py-2">平均账单额</th>
            <th class="border border-gray-300 px-4 py-2">平均还款额</th>
            <th class="border border-gray-300 px-4 py-2">逾期客户平均账单</th>
            <th class="border border-gray-300 px-4 py-2">正常客户还款率</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in filteredTableData" :key="item.month" class="hover:bg-gray-50">
            <td class="border border-gray-300 px-4 py-2">{{ item.month }}</td>
            <td class="border border-gray-300 px-4 py-2">{{ formatNumber(item.avgBill) }}</td>
            <td class="border border-gray-300 px-4 py-2">{{ formatNumber(item.avgRepayment) }}</td>
            <td class="border border-gray-300 px-4 py-2">{{ formatNumber(item.avgOverdueBill) }}</td>
            <td class="border border-gray-300 px-4 py-2">{{ (item.repaymentRate * 100).toFixed(2) }}%</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch, nextTick } from 'vue'
import * as d3 from 'd3'

// 响应式数据
const loading = ref(true)
const error = ref(null)
const radarData = ref([])
const financialData = ref([])
const months = ref([])
const selectedMonth = ref('all')
const filteredTableData = ref([])

const chart = ref(null)

const colorScale = d3.scaleOrdinal()
  .range(["#5B8FF9", "#5AD8A6", "#5D7092", "#FF6B9D", "#C9CDD4", "#DDB27C"])

// 获取API数据
async function fetchData() {
  try {
    loading.value = true
    error.value = null
    
    const response = await fetch('http://zzh.hengtai119.cn/api/billandrepayment')
    
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    
    const apiData = await response.json()
    
    // 转换API数据为组件所需格式
    financialData.value = apiData.map(item => ({
      month: item.month,
      avgBill: item.averageBill,
      avgRepayment: item.averageRepay,
      avgOverdueBill: item.overdueAvgBill,
      repaymentRate: item.normalRepayRate
    }))
    
    // 生成雷达图数据
    radarData.value = apiData.map(item => ({
      className: `data-${item.month}`,
      axes: [
        { axis: "平均账单额", value: item.averageBill },
        { axis: "平均还款额", value: item.averageRepay },
        { axis: "逾期客户平均账单", value: item.overdueAvgBill },
        { axis: "正常客户还款率", value: item.normalRepayRate * 100 } // 转换为百分比显示
      ]
    }))
    
    // 生成月份列表
    months.value = apiData.map(item => item.month)
    
    // 初始化过滤数据
    filteredTableData.value = financialData.value
    
    // 更新颜色比例尺的域
    colorScale.domain(radarData.value.map(d => d.className))
    
    loading.value = false
    
    // 等待DOM更新后绘制图表
    await nextTick()
    createRadarChart(radarData.value)
    
  } catch (err) {
    console.error('Failed to fetch data:', err)
    error.value = `数据加载失败: ${err.message}`
    loading.value = false
  }
}

// 绘制雷达图
function createRadarChart(data) {
  if (!chart.value || !data.length) return

  d3.select(chart.value).selectAll("*").remove() // 清空

  const width = Math.min(1000, chart.value.clientWidth || 800)
  const height = 500
  const margin = { top: 50, right: 50, bottom: 50, left: 50 }
  const innerWidth = width - margin.left - margin.right
  const innerHeight = height - margin.top - margin.bottom
  const radius = Math.min(innerWidth, innerHeight) / 2

  const svg = d3.select(chart.value)
    .append("svg")
    .attr("width", width)
    .attr("height", height)
    .append("g")
    .attr("transform", `translate(${margin.left + innerWidth / 2},${margin.top + innerHeight / 2})`)

  const allAxes = data[0].axes.map(d => d.axis)
  const totalAxes = allAxes.length
  const angleSlice = (Math.PI * 2) / totalAxes

  // 动态计算最大值
  const maxValues = {}
  allAxes.forEach(axis => {
    const values = data.flatMap(d => d.axes.filter(a => a.axis === axis).map(a => a.value))
    const max = Math.max(...values)
    
    // 根据轴类型设置合适的最大值和刻度
    if (axis === "正常客户还款率") {
      maxValues[axis] = Math.ceil(max / 5) * 5 // 向上取整到5的倍数
    } else {
      maxValues[axis] = Math.ceil(max / 10000) * 10000 // 向上取整到万的倍数
    }
  })

  const axisConfig = {
    "平均账单额": { 
      max: maxValues["平均账单额"], 
      ticks: Array.from({length: 6}, (_, i) => (i + 1) * maxValues["平均账单额"] / 6)
    },
    "平均还款额": { 
      max: maxValues["平均还款额"], 
      ticks: Array.from({length: 5}, (_, i) => (i + 1) * maxValues["平均还款额"] / 5)
    },
    "逾期客户平均账单": { 
      max: maxValues["逾期客户平均账单"], 
      ticks: Array.from({length: 6}, (_, i) => (i + 1) * maxValues["逾期客户平均账单"] / 6)
    },
    "正常客户还款率": { 
      max: maxValues["正常客户还款率"], 
      ticks: Array.from({length: 5}, (_, i) => (i + 1) * maxValues["正常客户还款率"] / 5)
    }
  }

  // 创建比例尺
  const radialScales = {}
  allAxes.forEach(axis => {
    radialScales[axis] = d3.scaleLinear()
      .domain([0, axisConfig[axis].max])
      .range([0, radius])
  })

  // 网格圆
  svg.append("circle")
    .attr("r", radius)
    .attr("fill", "none")
    .attr("stroke", "#e2e8f0")
    .attr("stroke-width", 1)

  // 绘制第一个轴的刻度文本（避免刻度重复）
  axisConfig[allAxes[0]].ticks.forEach(tick => {
    const r = radialScales[allAxes[0]](tick)
    svg.append("text")
      .attr("x", r * Math.cos(angleSlice * 0 - Math.PI / 2))
      .attr("y", r * Math.sin(angleSlice * 0 - Math.PI / 2) - 5)
      .attr("text-anchor", "middle")
      .attr("fill", "#94a3b8")
      .attr("font-size", 10)
      .text(Math.round(tick) + (allAxes[0] === "正常客户还款率" ? "%" : ""))
  })

  // 轴线和轴标签
  allAxes.forEach((axis, i) => {
    const angle = angleSlice * i

    svg.append("line")
      .attr("x1", 0)
      .attr("y1", 0)
      .attr("x2", radialScales[axis](axisConfig[axis].max) * Math.cos(angle - Math.PI / 2))
      .attr("y2", radialScales[axis](axisConfig[axis].max) * Math.sin(angle - Math.PI / 2))
      .attr("stroke", "#cbd5e1")
      .attr("stroke-width", 1)

    svg.append("text")
      .attr("x", (radialScales[axis](axisConfig[axis].max) + 20) * Math.cos(angle - Math.PI / 2))
      .attr("y", (radialScales[axis](axisConfig[axis].max) + 20) * Math.sin(angle - Math.PI / 2))
      .attr("text-anchor", angle > Math.PI ? "end" : "start")
      .attr("fill", "#334155")
      .attr("font-weight", "bold")
      .text(axis)
 })

  // 画多边形
  const radarLine = d3.lineRadial()
    .radius(d => d.radius)
    .angle((d, i) => i * angleSlice)
    .curve(d3.curveLinearClosed)

  data.forEach(d => {
    const points = d.axes.map(axisData => ({
      axis: axisData.axis,
      radius: radialScales[axisData.axis](axisData.value)
    }))

    svg.append("path")
      .datum(points)
      .attr("d", radarLine)
      .attr("fill", colorScale(d.className))
      .attr("fill-opacity", 0.4)
      .attr("stroke", colorScale(d.className))
      .attr("stroke-width", 2)

    svg.selectAll(`.circle-point-${d.className}`)
      .data(points)
      .enter()
      .append("circle")
      .attr("class", `circle-point-${d.className}`)
      .attr("cx", (d, i) => d.radius * Math.cos(i * angleSlice - Math.PI / 2))
      .attr("cy", (d, i) => d.radius * Math.sin(i * angleSlice - Math.PI / 2))
      .attr("r", 4)
      .attr("fill", colorScale(d.className))
      .attr("stroke", "#fff")
      .attr("stroke-width", 1.5)
  })
}

function formatNumber(val) {
  return val.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 2 })
}

function handleMonthChange() {
  if (selectedMonth.value === "all") {
    filteredTableData.value = financialData.value
    createRadarChart(radarData.value)
  } else {
    filteredTableData.value = financialData.value.filter(d => d.month === selectedMonth.value)
    createRadarChart(radarData.value.filter(d => d.className === `data-${selectedMonth.value}`))
  }
}

onMounted(() => {
  fetchData()
})
</script>

<style scoped>
.custom-scrollbar {
  scrollbar-width: thin;
  scrollbar-color: #94a3b8 transparent;
}
.custom-scrollbar::-webkit-scrollbar {
  height: 8px;
  width: 8px;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background-color: #94a3b8;
  border-radius: 4px;
}
</style>
