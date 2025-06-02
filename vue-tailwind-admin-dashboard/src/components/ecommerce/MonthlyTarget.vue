<template>
  <div class="bg-white rounded-xl p-6 shadow-sm dark:bg-gray-900 relative">
    <h2 class="text-2xl font-semibold text-gray-800 dark:text-white mb-6">风险等级占比</h2>
    <div ref="svgRef" class="w-full h-[400px]" />
    <div ref="tooltip" class="tooltip hidden absolute px-2 py-1 rounded bg-gray-800 text-white text-sm pointer-events-none"></div>
  </div>
</template>

<script setup>
import * as d3 from 'd3'
import { onMounted, ref } from 'vue'

const svgRef = ref(null)

const data = [
  { label: '无逾期', value: 19931, color: '#10B981' },
  { label: '轻度逾期', value: 6772, color: '#F59E0B' },
  { label: '高风险', value: 3297, color: '#EF4444' },
]

onMounted(() => {
  const container = svgRef.value
  const tooltip = container.parentElement.querySelector('.tooltip')
  const width = container.offsetWidth
  const height = 400
  const radius = Math.min(width, height) / 2 - 20

  const svg = d3.select(container)
    .append('svg')
    .attr('width', width)
    .attr('height', height)
    .append('g')
    .attr('transform', `translate(${width / 2}, ${height / 2})`)

  const pie = d3.pie()
    .sort(null)
    .value(d => d.value)

  const arc = d3.arc()
    .innerRadius(radius * 0.6)  // 设为60%半径，变成圆环
    .outerRadius(radius)

  const arcs = svg.selectAll('arc')
    .data(pie(data))
    .enter()
    .append('g')

  arcs.append('path')
    .attr('d', arc)
    .attr('fill', d => d.data.color)
    .attr('stroke', 'white')
    .attr('stroke-width', 2)
    .on('mouseover', (event, d) => {
      tooltip.classList.remove('hidden')
      tooltip.textContent = `${d.data.label}: ${d.data.value} (${((d.data.value / d3.sum(data, d => d.value)) * 100).toFixed(1)}%)`
    })
    .on('mousemove', (event) => {
      const [x, y] = d3.pointer(event, container)
      const tooltipWidth = tooltip.offsetWidth
      const tooltipHeight = tooltip.offsetHeight
      let left = x + 10
      let top = y + 10
      if (left + tooltipWidth > width) {
        left = x - tooltipWidth - 10
      }
      if (top + tooltipHeight > height) {
        top = y - tooltipHeight - 10
      }
      tooltip.style.left = `${left}px`
      tooltip.style.top = `${top}px`
    })
    .on('mouseout', () => {
      tooltip.classList.add('hidden')
    })

  arcs.append('text')
    .attr('transform', d => `translate(${arc.centroid(d)})`)
    .attr('text-anchor', 'middle')
    .attr('alignment-baseline', 'middle')
    .attr('fill', 'white')
    .attr('font-size', '12px')
    .text(d => `${((d.data.value / d3.sum(data, d => d.value)) * 100).toFixed(1)}%`)

  // 图例
  const legend = svg.append('g')
    .attr('transform', `translate(${-width / 2 + 20}, ${-height / 2 + 20})`)

  data.forEach((d, i) => {
    const g = legend.append('g').attr('transform', `translate(0, ${i * 20})`)
    g.append('rect')
      .attr('width', 12)
      .attr('height', 12)
      .attr('fill', d.color)
    g.append('text')
      .attr('x', 18)
      .attr('y', 10)
      .attr('font-size', '12px')
      .attr('fill', '#4B5563')
      .text(d.label)
  })
})
</script>

<style scoped>
.tooltip {
  position: absolute;
  pointer-events: none;
  white-space: nowrap;
  transition: opacity 0.2s ease;
  opacity: 0.9;
  z-index: 10;
}
svg {
  max-width: 100%;
}
</style>

