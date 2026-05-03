<template>
  <div class="dashboard-container">
    <el-row :gutter="20" class="stats-row">
      <el-col :xs="12" :sm="12" :lg="6">
        <div class="stats-card card-green">
          <div class="card-title">总订单数</div>
          <div class="card-value">{{ statistics.total_orders || 0 }}</div>
        </div>
      </el-col>
      <el-col :xs="12" :sm="12" :lg="6">
        <div class="stats-card card-blue">
          <div class="card-title">总销售额</div>
          <div class="card-value">¥{{ statistics.total_sales || 0 }}</div>
        </div>
      </el-col>
      <el-col :xs="12" :sm="12" :lg="6">
        <div class="stats-card card-orange">
          <div class="card-title">商品总数</div>
          <div class="card-value">{{ statistics.total_products || 0 }}</div>
        </div>
      </el-col>
      <el-col :xs="12" :sm="12" :lg="6">
        <div class="stats-card">
          <div class="card-title">用户总数</div>
          <div class="card-value">{{ statistics.total_users || 0 }}</div>
        </div>
      </el-col>
    </el-row>

    <el-row :gutter="20" class="charts-row">
      <el-col :xs="24" :sm="24" :lg="12">
        <div class="page-container">
          <div class="page-header">
            <span class="page-title">订单统计</span>
          </div>
          <div ref="orderChart" class="chart-container"></div>
        </div>
      </el-col>
      <el-col :xs="24" :sm="24" :lg="12">
        <div class="page-container">
          <div class="page-header">
            <span class="page-title">销售额统计</span>
          </div>
          <div ref="salesChart" class="chart-container"></div>
        </div>
      </el-col>
    </el-row>

    <el-row :gutter="20" class="charts-row">
      <el-col :xs="24" :sm="24" :lg="12">
        <div class="page-container">
          <div class="page-header">
            <span class="page-title">热销商品 TOP5</span>
          </div>
          <el-table :data="hotProducts" style="width: 100%">
            <el-table-column prop="rank" label="排名" width="80">
              <template #default="scope">
                <el-tag :type="scope.row.rank <= 3 ? 'danger' : 'info'" effect="dark">
                  {{ scope.row.rank }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="name" label="商品名称" />
            <el-table-column prop="sales" label="销量" />
            <el-table-column prop="revenue" label="销售额">
              <template #default="scope">
                <span class="color-primary">¥{{ scope.row.revenue }}</span>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </el-col>
      <el-col :xs="24" :sm="24" :lg="12">
        <div class="page-container">
          <div class="page-header">
            <span class="page-title">商品分类占比</span>
          </div>
          <div ref="categoryChart" class="chart-container"></div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, onUnmounted } from 'vue'
import * as echarts from 'echarts'
import { get } from '@/api/request'
import { useUserStore } from '@/store/user'

const userStore = useUserStore()

const orderChart = ref(null)
const salesChart = ref(null)
const categoryChart = ref(null)

let orderChartInstance = null
let salesChartInstance = null
let categoryChartInstance = null

const statistics = reactive({
  total_orders: 0,
  total_sales: 0,
  total_products: 0,
  total_users: 0
})

const hotProducts = ref([
  { rank: 1, name: '有机蔬菜套餐', sales: 156, revenue: 12480 },
  { rank: 2, name: '新鲜水果礼盒', sales: 128, revenue: 9600 },
  { rank: 3, name: '农家土鸡蛋', sales: 95, revenue: 4750 },
  { rank: 4, name: '绿色大米', sales: 82, revenue: 3280 },
  { rank: 5, name: '蜂蜜礼盒', sales: 67, revenue: 4020 }
])

const fetchStatistics = async () => {
  try {
    let url = '/api/admin/statistics'
    if (userStore.isMerchant()) {
      url = '/api/merchant/statistics'
    }
    
    const res = await get(url)
    if (res.code === 200) {
      Object.assign(statistics, res.data)
    }
  } catch (error) {
    console.error('Fetch statistics error:', error)
  }
}

const initOrderChart = () => {
  if (!orderChart.value) return
  
  orderChartInstance = echarts.init(orderChart.value)
  
  const option = {
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'shadow'
      }
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'category',
      data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日']
    },
    yAxis: {
      type: 'value',
      name: '订单数'
    },
    series: [
      {
        name: '订单数',
        type: 'bar',
        data: [12, 19, 15, 22, 28, 35, 30],
        itemStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            { offset: 0, color: '#4caf50' },
            { offset: 1, color: '#8bc34a' }
          ])
        },
        barWidth: '40%'
      }
    ]
  }
  
  orderChartInstance.setOption(option)
}

const initSalesChart = () => {
  if (!salesChart.value) return
  
  salesChartInstance = echarts.init(salesChart.value)
  
  const option = {
    tooltip: {
      trigger: 'axis'
    },
    legend: {
      data: ['销售额'],
      top: 0
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: ['1月', '2月', '3月', '4月', '5月', '6月']
    },
    yAxis: {
      type: 'value',
      name: '元'
    },
    series: [
      {
        name: '销售额',
        type: 'line',
        smooth: true,
        data: [12000, 18000, 22000, 28000, 35000, 42000],
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            { offset: 0, color: 'rgba(76, 175, 80, 0.3)' },
            { offset: 1, color: 'rgba(76, 175, 80, 0.05)' }
          ])
        },
        lineStyle: {
          color: '#4caf50',
          width: 2
        },
        itemStyle: {
          color: '#4caf50'
        }
      }
    ]
  }
  
  salesChartInstance.setOption(option)
}

const initCategoryChart = () => {
  if (!categoryChart.value) return
  
  categoryChartInstance = echarts.init(categoryChart.value)
  
  const option = {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      orient: 'vertical',
      right: '5%',
      top: 'center'
    },
    series: [
      {
        name: '商品分类',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 10,
          borderColor: '#fff',
          borderWidth: 2
        },
        label: {
          show: false,
          position: 'center'
        },
        emphasis: {
          label: {
            show: true,
            fontSize: 16,
            fontWeight: 'bold'
          }
        },
        labelLine: {
          show: false
        },
        data: [
          { value: 1048, name: '蔬菜', itemStyle: { color: '#4caf50' } },
          { value: 735, name: '水果', itemStyle: { color: '#ff9800' } },
          { value: 580, name: '粮油', itemStyle: { color: '#2196f3' } },
          { value: 484, name: '肉禽', itemStyle: { color: '#f44336' } },
          { value: 300, name: '其他', itemStyle: { color: '#9c27b0' } }
        ]
      }
    ]
  }
  
  categoryChartInstance.setOption(option)
}

const handleResize = () => {
  orderChartInstance?.resize()
  salesChartInstance?.resize()
  categoryChartInstance?.resize()
}

onMounted(() => {
  fetchStatistics()
  initOrderChart()
  initSalesChart()
  initCategoryChart()
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
  orderChartInstance?.dispose()
  salesChartInstance?.dispose()
  categoryChartInstance?.dispose()
})
</script>

<style lang="scss" scoped>
.dashboard-container {
  .stats-row {
    margin-bottom: 20px;
  }

  .charts-row {
    margin-bottom: 20px;
  }
}
</style>
