<template>
  <AdminLayout>
    <PageBreadcrumb :pageTitle="currentPageTitle" />
    <div class="space-y-5 sm:space-y-6">
      <ComponentCard title="LimitBal Defaults Chart">
        <LineChartOne />
      </ComponentCard>
    </div>
  </AdminLayout>
</template>

<script setup>
import LineChartOne from "@/components/charts/DefaultChart/LimitBalDefaultsChart.vue";
import ComponentCard from "@/components/common/ComponentCard.vue";
import PageBreadcrumb from "@/components/common/PageBreadcrumb.vue";
import AdminLayout from "@/components/layout/AdminLayout.vue";
import { ref } from "vue";
const currentPageTitle = ref("LimitBal Defaults Chart");
</script>
