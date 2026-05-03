<template>
  <div class="recharge-container">
    <van-nav-bar
      title="账户充值"
      left-arrow
      @click-left="onClickLeft"
    />

    <div class="balance-section">
      <div class="balance-label">账户余额</div>
      <div class="balance-value">
        <span class="symbol">¥</span>
        <span class="amount">{{ balance }}</span>
      </div>
    </div>

    <div class="amount-section">
      <div class="section-title">选择充值金额</div>
      <van-grid :column-num="3" gutter="10">
        <van-grid-item
          v-for="(amount, index) in amountOptions"
          :key="index"
          @click="selectAmount(amount)"
        >
          <div
            class="amount-item"
            :class="{ 'active': selectedAmount === amount }"
          >
            <div class="amount-value">¥{{ amount }}</div>
          </div>
        </van-grid-item>
      </van-grid>

      <van-field
        v-model="customAmount"
        type="number"
        placeholder="或输入自定义金额"
        class="custom-amount"
        @input="onCustomAmountInput"
      />
    </div>

    <div class="pay-section">
      <div class="section-title">支付方式</div>
      <van-radio-group v-model="selectedPayType">
        <van-cell-group inset>
          <van-cell clickable @click="selectedPayType = 'wechat'">
            <template #icon>
              <van-icon name="wechat" size="24" color="#07c160" />
            </template>
            <div class="pay-text">微信支付</div>
            <template #right-icon>
              <van-radio name="wechat" />
            </template>
          </van-cell>
          <van-cell clickable @click="selectedPayType = 'alipay'">
            <template #icon>
              <van-icon name="alipay" size="24" color="#1677ff" />
            </template>
            <div class="pay-text">支付宝</div>
            <template #right-icon>
              <van-radio name="alipay" />
            </template>
          </van-cell>
        </van-cell-group>
      </van-radio-group>
    </div>

    <div class="submit-section">
      <van-button type="primary" round block :loading="loading" @click="handleRecharge">
        确认充值 ¥{{ totalAmount }}
      </van-button>
    </div>

    <van-notice-bar left-icon="warning-o" mode="closeable">
      充值金额将立即到账，请确认支付金额后再进行支付
    </van-notice-bar>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { showToast, showSuccessToast } from 'vant'

const router = useRouter()

const balance = ref('128.50')
const selectedAmount = ref(0)
const customAmount = ref('')
const selectedPayType = ref('wechat')
const loading = ref(false)

const amountOptions = ref([10, 20, 50, 100, 200, 500])

const totalAmount = computed(() => {
  return customAmount.value ? parseFloat(customAmount.value) || 0 : selectedAmount.value
})

const onClickLeft = () => {
  router.back()
}

const selectAmount = (amount) => {
  selectedAmount.value = amount
  customAmount.value = ''
}

const onCustomAmountInput = () => {
  if (customAmount.value) {
    selectedAmount.value = 0
  }
}

const handleRecharge = async () => {
  if (totalAmount.value <= 0) {
    showToast('请选择或输入充值金额')
    return
  }
  if (totalAmount.value < 1) {
    showToast('充值金额不能小于1元')
    return
  }

  loading.value = true
  setTimeout(() => {
    loading.value = false
    const newBalance = (parseFloat(balance.value) + totalAmount.value).toFixed(2)
    balance.value = newBalance
    showSuccessToast(`充值成功，到账 ¥${totalAmount.value}`)
    setTimeout(() => {
      router.back()
    }, 1500)
  }, 1500)
}
</script>

<style lang="scss" scoped>
.recharge-container {
  background-color: #f5f5f5;
  min-height: 100vh;
  padding-bottom: 20px;
}

.balance-section {
  background: linear-gradient(135deg, #4caf50 0%, #8bc34a 100%);
  padding: 30px 20px;
  text-align: center;
  color: #fff;

  .balance-label {
    font-size: 14px;
    opacity: 0.9;
    margin-bottom: 10px;
  }

  .balance-value {
    .symbol {
      font-size: 20px;
      margin-right: 4px;
    }

    .amount {
      font-size: 36px;
      font-weight: bold;
    }
  }
}

.amount-section,
.pay-section {
  padding: 15px;
  background-color: #fff;
  margin-top: 10px;
}

.section-title {
  font-size: 15px;
  font-weight: bold;
  color: #333;
  margin-bottom: 15px;
}

.amount-item {
  padding: 15px 0;
  border: 1px solid #eee;
  border-radius: 8px;
  text-align: center;

  &.active {
    border-color: #4caf50;
    background-color: rgba(76, 175, 80, 0.1);

    .amount-value {
      color: #4caf50;
    }
  }
}

.amount-value {
  font-size: 16px;
  font-weight: bold;
  color: #333;
}

.custom-amount {
  margin-top: 15px;

  :deep(.van-field__control) {
    border: 1px solid #eee;
    border-radius: 8px;
    padding: 10px 15px;
    background-color: #fafafa;
  }
}

.pay-text {
  font-size: 15px;
  color: #333;
}

.submit-section {
  padding: 20px;
}

:deep(.van-notice-bar) {
  margin-top: 10px;
}
</style>
