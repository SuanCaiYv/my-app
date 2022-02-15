<template>
    <div class="title">
        <input type="text" class="input" v-model="title" placeholder="文章标题">
    </div>
</template>

<script setup lang="ts">
import {inject, ref} from "vue"

const name = ref<string>("Title")

const title = inject("title")
</script>

<style scoped>
.title {
    width: 100%;
    height: 100%;
    grid-area: title;
}

.input {
    height: 100%;
    width: 100%;
    border: none;
    padding: 0 0 0 4px;
    font-size: 1.6rem;
    font-weight: bolder;
    outline: none;
}
</style>