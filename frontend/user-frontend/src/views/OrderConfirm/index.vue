<template>
  <div class="order-confirm-container">
    <van-nav-bar
      title="确认订单"
      left-arrow
      @click-left="onClickLeft"
    />

    <van-cell-group inset class="address-section">
      <van-cell
        is-link
        @click="showAddressPicker = true"
      >
        <template #default>
          <div v-if="selectedAddress" class="address-info">
            <div class="address-header">
              <span class="name">{{ selectedAddress.name }}</span>
              <span class="phone">{{ selectedAddress.phone }}</span>
              <van-tag v-if="selectedAddress.is_default" type="primary" size="small">默认</van-tag>
            </div>
            <div class="address-detail text-gray">{{ selectedAddress.province }}{{ selectedAddress.city }}{{ selectedAddress.district }}{{ selectedAddress.detail }}</div>
          </div>
          <div v-else class="select-addr">
            <span class="text-gray">请选择收货地址</span>
            <van-icon name="arrow" />
          </div>
        </template>
      </van-cell>
    </van-cell-group>

    <van-cell-group inset class="product-section">
      <div class="product-item flex" v-for="item in cartItems" :key="item.id">
        <img :src="item.image || getDefaultImage()" class="product-image" />
        <div class="product-info flex-1">
          <div class="product-name text-ellipsis-2">{{ item.name }}</div>
          <div class="product-price-row flex-between">
            <span class="price">¥{{ item.price }}</span>
            <span class="quantity">x{{ item.quantity }}</span>
          </div>
        </div>
      </div>
    </van-cell-group>

    <van-cell-group inset class="info-section">
      <van-cell title="配送方式">
        <template #right-icon>
          <span>快递配送</span>
        </template>
      </van-cell>
      <van-cell title="支付方式" is-link @click="showPayTypePicker = true">
        <template #right-icon>
          <span>{{ payTypeText }}</span>
        </template>
      </van-cell>
      <van-cell title="订单备注">
        <template #right-icon>
          <van-field
            v-model="remark"
            placeholder="请输入备注信息（选填）"
            :border="false"
            class="remark-input"
          />
        </template>
      </van-cell>
    </van-cell-group>

    <van-submit-bar
      :price="totalPrice * 100"
      :button-text="'提交订单'"
      @submit="handleSubmit"
    >
      <div class="price-info">
        <div class="price-row">
          <span class="label">商品金额</span>
          <span class="value">¥{{ goodsAmount }}</span>
        </div>
        <div class="price-row" v-if="shippingFee > 0">
          <span class="label">运费</span>
          <span class="value">¥{{ shippingFee }}</span>
        </div>
        <div class="price-row total">
          <span class="label">实付金额</span>
          <span class="value color-danger font-bold">¥{{ totalPrice }}</span>
        </div>
      </div>
    </van-submit-bar>

    <van-popup v-model:show="showAddressPicker" position="bottom" round>
      <div class="picker-header flex-between">
        <span class="picker-title">选择收货地址</span>
        <van-button type="primary" size="small" @click="goAddressAdd">新增地址</van-button>
      </div>
      <van-radio-group v-model="selectedAddressId">
        <van-cell
          v-for="address in addressList"
          :key="address.id"
          clickable
          @click="selectAddress(address)"
        >
          <template #default>
            <div class="address-option">
              <div class="address-option-header">
                <van-radio :name="address.id" />
                <span class="name">{{ address.name }}</span>
                <span class="phone">{{ address.phone }}</span>
                <van-tag v-if="address.is_default" type="primary" size="small">默认</van-tag>
              </div>
              <div class="address-option-detail text-gray">{{ address.province }}{{ address.city }}{{ address.district }}{{ address.detail }}</div>
            </div>
          </template>
        </van-cell>
      </van-radio-group>
    </van-popup>

    <van-action-sheet
      v-model:show="showPayTypePicker"
      :actions="payTypeActions"
      cancel-text="取消"
      @select="onSelectPayType"
    />
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'

const router = useRouter()

const getDefaultImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=fresh%20vegetables%20and%20fruits&image_size=square'
}

const selectedAddress = ref(null)
const selectedAddressId = ref(null)
const remark = ref('')
const showAddressPicker = ref(false)
const showPayTypePicker = ref(false)
const payType = ref('balance')

const addressList = ref([
  { id: 1, name: '张三', phone: '13800138000', province: '北京市', city: '北京市', district: '朝阳区', detail: '某某街道某某小区1号楼101室', is_default: true },
  { id: 2, name: '李四', phone: '13800138001', province: '上海市', city: '上海市', district: '浦东新区', detail: '某某路某某大厦A座505室', is_default: false }
])

const cartItems = ref([
  { id: 1, name: '有机蔬菜套餐', price: 68.00, quantity: 2, image: '' },
  { id: 2, name: '新鲜水果礼盒', price: 22.50, quantity: 1, image: '' }
])

const goodsAmount = computed(() => {
  return cartItems.value.reduce((sum, item) => sum + item.price * item.quantity, 0)
})

const shippingFee = computed(() => {
  return goodsAmount.value >= 99 ? 0 : 10
})

const totalPrice = computed(() => {
  return goodsAmount.value + shippingFee.value
})

const payTypeText = computed(() => {
  const texts = {
    balance: '余额支付',
    wechat: '微信支付',
    alipay: '支付宝'
  }
  return texts[payType.value] || '请选择'
})

const payTypeActions = ref([
  { name: '余额支付', value: 'balance' },
  { name: '微信支付', value: 'wechat' },
  { name: '支付宝', value: 'alipay' }
])

const onClickLeft = () => {
  router.back()
}

const goAddressAdd = () => {
  showAddressPicker.value = false
  router.push('/address/add')
}

const selectAddress = (address) => {
  selectedAddress.value = address
  selectedAddressId.value = address.id
  showAddressPicker.value = false
}

const onSelectPayType = (action) => {
  payType.value = action.value
}

const handleSubmit = () => {
  if (!selectedAddress.value) {
    showToast('请选择收货地址')
    return
  }
  showToast('订单提交成功')
  setTimeout(() => {
    router.push('/order')
  }, 1000)
}
</script>

<style lang="scss" scoped>
.order-confirm-container {
  background-color: #f5f5f5;
  min-height: 100vh;
  padding-bottom: 60px;
}

.address-section,
.product-section,
.info-section {
  margin-top: 10px;
}

.address-info {
  .address-header {
    margin-bottom: 8px;
    display: flex;
    align-items: center;

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
}

.select-addr {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
}

.product-item {
  padding: 12px 15px;
  border-bottom: 1px solid #f5f5f5;

  &:last-child {
    border-bottom: none;
  }
}

.product-image {
  width: 80px;
  height: 80px;
  border-radius: 4px;
  object-fit: cover;
}

.product-info {
  margin-left: 12px;
}

.product-name {
  font-size: 14px;
  color: #333;
  margin-bottom: 10px;
  line-height: 1.4;
}

.product-price-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.price {
  font-size: 14px;
  color: #f44336;
  font-weight: bold;
}

.quantity {
  font-size: 13px;
  color: #999;
}

.remark-input {
  text-align: right;
}

:deep(.van-submit-bar__text) {
  flex: 1;
  padding: 0;
}

.price-info {
  .price-row {
    display: flex;
    justify-content: space-between;
    font-size: 13px;
    margin-bottom: 5px;

    &.total {
      font-size: 14px;
      margin-top: 8px;
      padding-top: 8px;
      border-top: 1px solid #f5f5f5;
    }

    .label {
      color: #666;
    }

    .value {
      color: #333;
    }
  }
}

.picker-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px;
  border-bottom: 1px solid #f5f5f5;
}

.picker-title {
  font-size: 16px;
  font-weight: bold;
}

.address-option {
  .address-option-header {
    display: flex;
    align-items: center;
    margin-bottom: 8px;

    .name {
      margin-left: 10px;
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

  .address-option-detail {
    margin-left: 32px;
    font-size: 14px;
    color: #666;
    line-height: 1.5;
  }
}

.text-gray {
  color: #999;
}

.color-danger {
  color: #f44336;
}

.font-bold {
  font-weight: bold;
}
</style>
