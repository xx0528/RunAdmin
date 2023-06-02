<template>
  <div>
    <div class="gva-form-box">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="工单名:" prop="orderName">
          <el-input v-model="formData.orderName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="落地页ID:" prop="pageId">
          <el-input v-model="formData.pageId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="进粉限制:" prop="maxEnterNum">
          <el-input v-model.number="formData.maxEnterNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="平均进粉:" prop="eachEnterNum">
          <el-input v-model.number="formData.eachEnterNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'RunOrder'
}
</script>

<script setup>
import {
  createRunOrder,
  updateRunOrder,
  findRunOrder
} from '@/api/runOrder'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  orderName: '',
  pageId: '',
  maxEnterNum: 0,
  eachEnterNum: 0,
})
// 验证规则
const rule = reactive({
  orderName: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  pageId: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  maxEnterNum: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  eachEnterNum: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findRunOrder({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.rerunOrder
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate(async(valid) => {
        if (!valid) return
        let res
        switch (type.value) {
          case 'create':
            res = await createRunOrder(formData.value)
            break
          case 'update':
            res = await updateRunOrder(formData.value)
            break
          default:
            res = await createRunOrder(formData.value)
            break
        }
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '创建/更改成功'
          })
        }
      })
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style>
</style>
