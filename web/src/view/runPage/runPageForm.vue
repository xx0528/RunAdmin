<template>
  <div>
    <div class="gva-form-box">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名字:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="链接:" prop="url">
          <el-input v-model="formData.url" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="国家:" prop="country">
          <el-input v-model="formData.country" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户数:" prop="userNum">
          <el-input v-model.number="formData.userNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户id:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:" prop="state">
          <el-select v-model="formData.state" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in stateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  name: 'RunPage'
}
</script>

<script setup>
import {
  createRunPage,
  updateRunPage,
  findRunPage
} from '@/api/runPage'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const stateOptions = ref([])
const formData = ref({
  name: '',
  url: '',
  country: '',
  pageId: '',
  userNum: 0,
  userId: 0,
  state: undefined,
})
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findRunPage({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.rerunPage
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
  stateOptions.value = await getDictFunc('state')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate(async(valid) => {
        if (!valid) return
        let res
        switch (type.value) {
          case 'create':
            res = await createRunPage(formData.value)
            break
          case 'update':
            res = await updateRunPage(formData.value)
            break
          default:
            res = await createRunPage(formData.value)
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
