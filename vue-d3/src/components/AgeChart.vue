<template>
  <div>
    <h2>年龄分布柱状图</h2>
    <div ref="chart"></div>
  </div>
</template>

<script>
import * as d3 from "d3";
import axios from "axios";

export default {
  name: "AgeChart",
  mounted() {
    this.fetchData();
  },
  methods: {
    async fetchData() {
      try {
        const res = await axios.get("http://zzh.hengtai119.cn/api/ageinfo");
        if (res.data.code === 200 && Array.isArray(res.data.data)) {
          this.renderChart(res.data.data);
        } else {
          console.error("数据格式不正确：", res.data);
        }
      } catch (error) {
        console.error("请求失败：", error);
      }
    },
    renderChart(data) {
      const margin = { top: 20, right: 30, bottom: 40, left: 40 };
      const width = 600 - margin.left - margin.right;
      const height = 400 - margin.top - margin.bottom;

      const svg = d3.select(this.$refs.chart)
        .append("svg")
        .attr("width", width + margin.left + margin.right)
        .attr("height", height + margin.top + margin.bottom)
        .append("g")
        .attr("transform", `translate(${margin.left},${margin.top})`);

      // 处理数据
      const x = d3.scaleBand()
        .domain(data.map(d => d.Age.toString()))
        .range([0, width])
        .padding(0.1);

      const y = d3.scaleLinear()
        .domain([0, d3.max(data, d => d.Count)])
        .nice()
        .range([height, 0]);

      svg.append("g")
        .selectAll("rect")
        .data(data)
        .enter().append("rect")
        .attr("x", d => x(d.Age.toString()))
        .attr("y", d => y(d.Count))
        .attr("width", x.bandwidth())
        .attr("height", d => height - y(d.Count))
        .attr("fill", "#69b3a2");

      // 添加 X 轴
      svg.append("g")
        .attr("transform", `translate(0,${height})`)
        .call(d3.axisBottom(x));

      // 添加 Y 轴
      svg.append("g")
        .call(d3.axisLeft(y));
    }
  }
};
</script>

<style scoped>
svg {
  font-family: sans-serif;
}
</style>

