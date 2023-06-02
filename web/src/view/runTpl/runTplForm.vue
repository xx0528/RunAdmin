<template>
  <div>
    <div class="gva-form-box">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="模板页面ID:" prop="tplId">
          <el-input v-model="formData.tplId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="模板网页名:" prop="tplName">
          <el-input v-model="formData.tplName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="点击链接:" prop="clickUrl">
          <el-input v-model="formData.clickUrl" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="按钮描述:" prop="clickDesc">
          <el-input v-model="formData.clickDesc" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="弹窗描述:" prop="dialogDesc">
          <el-input v-model="formData.dialogDesc" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述1:" prop="text1">
          <el-input v-model="formData.text1" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述2:" prop="text2">
          <el-input v-model="formData.text2" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述3:" prop="text3">
          <el-input v-model="formData.text3" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述4:" prop="text4">
          <el-input v-model="formData.text4" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述5:" prop="text5">
          <el-input v-model="formData.text5" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述6:" prop="text6">
          <el-input v-model="formData.text6" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述7:" prop="text7">
          <el-input v-model="formData.text7" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述8:" prop="text8">
          <el-input v-model="formData.text8" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述9:" prop="text9">
          <el-input v-model="formData.text9" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述10:" prop="text10">
          <el-input v-model="formData.text10" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述11:" prop="text11">
          <el-input v-model="formData.text11" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述12:" prop="text12">
          <el-input v-model="formData.text12" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述13:" prop="text13">
          <el-input v-model="formData.text13" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述14:" prop="text14">
          <el-input v-model="formData.text14" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述15:" prop="text15">
          <el-input v-model="formData.text15" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述16:" prop="text16">
          <el-input v-model="formData.text16" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述17:" prop="text17">
          <el-input v-model="formData.text17" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述18:" prop="text18">
          <el-input v-model="formData.text18" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述19:" prop="text19">
          <el-input v-model="formData.text19" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述20:" prop="text20">
          <el-input v-model="formData.text20" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图片1:" prop="picName1">
          <el-input v-model="formData.picName1" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图片2:" prop="picName2">
          <el-input v-model="formData.picName2" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图片3:" prop="picName3">
          <el-input v-model="formData.picName3" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图片4:" prop="picName4">
          <el-input v-model="formData.picName4" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图片5:" prop="picName5">
          <el-input v-model="formData.picName5" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图片6:" prop="picName6">
          <el-input v-model="formData.picName6" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图片7:" prop="picName7">
          <el-input v-model="formData.picName7" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图片8:" prop="picName8">
          <el-input v-model="formData.picName8" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图片9:" prop="picName9">
          <el-input v-model="formData.picName9" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图片10:" prop="picName10">
          <el-input v-model="formData.picName10" :clearable="true" placeholder="请输入" />
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
  name: 'RunTpl'
}
</script>

<script setup>
import {
  createRunTpl,
  updateRunTpl,
  findRunTpl
} from '@/api/runTpl'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  tplId: '',
  tplName: '',
  clickUrl: '',
  clickDesc: '',
  dialogDesc: '',
  text1: '',
  text2: '',
  text3: '',
  text4: '',
  text5: '',
  text6: '',
  text7: '',
  text8: '',
  text9: '',
  text10: '',
  text11: '',
  text12: '',
  text13: '',
  text14: '',
  text15: '',
  text16: '',
  text17: '',
  text18: '',
  text19: '',
  text20: '',
  picName1: '',
  picName2: '',
  picName3: '',
  picName4: '',
  picName5: '',
  picName6: '',
  picName7: '',
  picName8: '',
  picName9: '',
  picName10: '',
})
// 验证规则
const rule = reactive({
  tplName: [{
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
    const res = await findRunTpl({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.rerunTpl
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
            res = await createRunTpl(formData.value)
            break
          case 'update':
            res = await updateRunTpl(formData.value)
            break
          default:
            res = await createRunTpl(formData.value)
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
