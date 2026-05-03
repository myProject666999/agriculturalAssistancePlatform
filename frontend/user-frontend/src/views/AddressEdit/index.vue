<template>
  <div class="address-edit-container">
    <van-nav-bar
      :title="isEdit ? '编辑地址' : '添加地址'"
      left-arrow
      @click-left="onClickLeft"
    />

    <van-form @submit="onSubmit">
      <van-cell-group inset>
        <van-field
          v-model="form.name"
          label="收货人"
          placeholder="请输入收货人姓名"
          :rules="[{ required: true, message: '请填写收货人姓名' }]"
        />
        <van-field
          v-model="form.phone"
          label="手机号"
          placeholder="请输入手机号"
          :rules="[{ required: true, message: '请填写手机号' }, { pattern: /^1[3-9]\d{9}$/, message: '手机号格式错误' }]"
        />
        <van-field
          v-model="areaText"
          is-link
          readonly
          label="所在地区"
          placeholder="请选择省/市/区"
          @click="showAreaPicker = true"
        />
        <van-field
          v-model="form.detail"
          label="详细地址"
          type="textarea"
          placeholder="请输入详细地址（街道、门牌号等）"
          maxlength="200"
          show-word-limit
          :rules="[{ required: true, message: '请填写详细地址' }]"
        />
      </van-cell-group>

      <van-cell-group inset class="switch-section">
        <van-cell title="设为默认地址">
          <template #right-icon>
            <van-switch
              v-model="form.is_default"
              active-color="#4caf50"
            />
          </template>
        </van-cell>
      </van-cell-group>

      <div class="form-actions">
        <van-button type="primary" round block :loading="loading" type="submit">
          保存
        </van-button>
      </div>
    </van-form>

    <van-popup v-model:show="showAreaPicker" position="bottom">
      <van-area
        :area-list="areaList"
        @confirm="onAreaConfirm"
        @cancel="showAreaPicker = false"
      />
    </van-popup>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { showToast } from 'vant'

const router = useRouter()
const route = useRoute()

const isEdit = computed(() => !!route.params.id)
const loading = ref(false)
const showAreaPicker = ref(false)

const form = reactive({
  name: '',
  phone: '',
  province: '',
  city: '',
  district: '',
  detail: '',
  is_default: false
})

const areaText = computed(() => {
  return [form.province, form.city, form.district].filter(Boolean).join(' / ')
})

const areaList = ref({
  province_list: {
    110000: '北京市',
    310000: '上海市',
    440000: '广东省',
    330000: '浙江省',
    320000: '江苏省'
  },
  city_list: {
    110100: '北京市',
    310100: '上海市',
    440100: '广州市',
    440300: '深圳市',
    330100: '杭州市',
    320100: '南京市'
  },
  county_list: {
    110101: '东城区',
    110105: '朝阳区',
    110106: '丰台区',
    310101: '黄浦区',
    310104: '徐汇区',
    440103: '荔湾区',
    440104: '越秀区',
    330102: '上城区',
    320102: '玄武区'
  }
})

const onClickLeft = () => {
  router.back()
}

const onAreaConfirm = ({ selectedOptions }) => {
  const [province, city, county] = selectedOptions
  form.province = province?.text || ''
  form.city = city?.text || ''
  form.district = county?.text || ''
  showAreaPicker.value = false
}

const onSubmit = async () => {
  if (!form.name.trim()) {
    showToast('请输入收货人姓名')
    return
  }
  if (!form.phone.trim()) {
    showToast('请输入手机号')
    return
  }
  if (!/^1[3-9]\d{9}$/.test(form.phone.trim())) {
    showToast('手机号格式错误')
    return
  }
  if (!form.detail.trim()) {
    showToast('请输入详细地址')
    return
  }

  loading.value = true
  setTimeout(() => {
    loading.value = false
    showToast(isEdit.value ? '修改成功' : '添加成功')
    setTimeout(() => {
      router.back()
    }, 1000)
  }, 1000)
}
</script>

<style lang="scss" scoped>
.address-edit-container {
  background-color: #f5f5f5;
  min-height: 100vh;
}

.switch-section {
  margin-top: 10px;
}

.form-actions {
  padding: 20px;
}
</style>
