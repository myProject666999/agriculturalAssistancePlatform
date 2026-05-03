<template>
  <div class="address-container">
    <van-nav-bar
      title="收货地址"
      left-arrow
      @click-left="onClickLeft"
      :fixed="true"
      :placeholder="true"
    >
      <template #right>
        <van-icon name="plus" size="20" @click="goAdd" />
      </template>
    </van-nav-bar>

    <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
      <van-list
        v-model:loading="loading"
        :finished="finished"
        finished-text="没有更多了"
        @load="onLoad"
      >
        <van-radio-group v-model="selectedId">
          <div
            v-for="address in addressList"
            :key="address.id"
            class="address-item"
          >
            <van-radio
              :name="address.id"
              v-if="isSelectMode"
              @click="selectAddress(address)"
            />
            <div class="address-content" :class="{ 'clickable': !isSelectMode }" @click="handleEdit(address)">
              <div class="address-header">
                <span class="name">{{ address.name }}</span>
                <span class="phone">{{ address.phone }}</span>
                <van-tag v-if="address.is_default" type="primary" size="small">默认</van-tag>
              </div>
              <div class="address-detail text-gray">{{ address.province }}{{ address.city }}{{ address.district }}{{ address.detail }}</div>
            </div>
            <div class="address-actions">
              <van-switch
                v-model="address.is_default"
                size="24"
                active-color="#4caf50"
                @change="setDefault(address)"
              />
              <van-icon name="edit" size="18" color="#666" @click="goEdit(address.id)" />
              <van-icon name="delete" size="18" color="#999" @click="handleDelete(address)" />
            </div>
          </div>
        </van-radio-group>

        <van-empty v-if="addressList.length === 0 && !loading" description="暂无收货地址" />
      </van-list>
    </van-pull-refresh>

    <div class="bottom-actions" v-if="!isSelectMode">
      <van-button type="primary" round block @click="goAdd">
        <van-icon name="plus" />
        添加收货地址
      </van-button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { showConfirmDialog, showToast } from 'vant'

const router = useRouter()
const route = useRoute()

const refreshing = ref(false)
const loading = ref(false)
const finished = ref(false)
const selectedId = ref(route.query.selectedId || null)
const isSelectMode = ref(!!route.query.from)

const addressList = ref([
  { id: 1, name: '张三', phone: '13800138000', province: '北京市', city: '北京市', district: '朝阳区', detail: '某某街道某某小区1号楼101室', is_default: true },
  { id: 2, name: '李四', phone: '13800138001', province: '上海市', city: '上海市', district: '浦东新区', detail: '某某路某某大厦A座505室', is_default: false }
])

const onClickLeft = () => {
  router.back()
}

const goAdd = () => {
  router.push('/address/add')
}

const goEdit = (id) => {
  router.push(`/address/edit/${id}`)
}

const handleEdit = (address) => {
  if (!isSelectMode.value) {
    goEdit(address.id)
  }
}

const selectAddress = (address) => {
  if (route.query.from === 'order') {
    router.back()
  }
}

const setDefault = (address) => {
  addressList.value.forEach(item => {
    item.is_default = item.id === address.id
  })
  showToast('已设为默认地址')
}

const handleDelete = (address) => {
  showConfirmDialog({
    title: '提示',
    message: '确定要删除该地址吗？'
  }).then(() => {
    addressList.value = addressList.value.filter(item => item.id !== address.id)
    showToast('已删除')
  }).catch(() => {})
}

const onLoad = () => {
  setTimeout(() => {
    if (refreshing.value) {
      refreshing.value = false
    }
    loading.value = false
    finished.value = true
  }, 1000)
}

const onRefresh = () => {
  onLoad()
}
</script>

<style lang="scss" scoped>
.address-container {
  background-color: #f5f5f5;
  min-height: 100vh;
  padding-bottom: 80px;
}

.address-item {
  display: flex;
  align-items: center;
  background-color: #fff;
  padding: 15px;
  margin-bottom: 10px;

  .van-radio {
    margin-right: 12px;
  }
}

.address-content {
  flex: 1;

  &.clickable {
    cursor: pointer;
  }
}

.address-header {
  display: flex;
  align-items: center;
  margin-bottom: 8px;

  .name {
    font-size: 15px;
    font-weight: bold;
    color: #333;
  }

  .phone {
    margin-left: 15px;
    font-size: 14px;
    color: #666;
  }

  .van-tag {
    margin-left: 10px;
  }
}

.address-detail {
  font-size: 14px;
  color: #666;
  line-height: 1.5;
}

.address-actions {
  display: flex;
  align-items: center;
  gap: 15px;
  margin-left: 10px;
}

.bottom-actions {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 15px;
  background-color: #fff;
  box-shadow: 0 -2px 10px rgba(0, 0, 0, 0.05);
}

.text-gray {
  color: #999;
}
</style>
