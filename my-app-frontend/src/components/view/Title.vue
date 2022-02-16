<template>
    <div class="title">{{props.value}}</div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<string>("Title")
const props = defineProps({
    value: String,
})
</script>

<style scoped>
.title {
    width: 100%;
    height: 60px;
    font-size: 1.4rem;
    border-bottom: 2px solid black;
    box-sizing: border-box;
    font-weight: bolder;
    text-align: left;
}
</style>