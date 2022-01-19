<template>
    <div class="tagList">
        <!--占位符，勿动！-->
        <div style="margin-bottom: 50px;box-sizing: border-box"></div>
        <div v-for="t in tagList">
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
tagList.push("bbb")
tagList.push("ccc")
</script>

<style scoped>
.tagList {
    width: 100%;
    height: 100%;
    border: 1px solid silver;
    box-sizing: border-box;
}
</style>