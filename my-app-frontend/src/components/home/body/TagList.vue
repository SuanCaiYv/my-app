<template>
    <div class="tagList">
        <!--占位符，勿动！-->
        <div style="margin-top: 25px;box-sizing: border-box"></div>
        <div v-for="t in tagList" style="display: inline-block">
            <Tag :value="t"></Tag>
        </div>
    </div>
</template>

<script setup lang="ts">
import {reactive, ref} from "vue"
import Tag from "./Tag.vue";

const name = ref<String>("TagList")

const tagList = reactive<Array<String>>([])

tagList.push("aaa")
tagList.push("bbbbbb")
tagList.push("ccc")
tagList.push("dddddddd")
tagList.push("ee")
</script>

<style scoped>
.tagList {
    width: 100%;
    height: 100%;
    border: 2px solid ;
    box-sizing: border-box;
    border-radius: 20px;
    background-color: ivory;
}
</style>