<template>
    <div class="ph4"></div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("PH4")
</script>

<style scoped>
.ph4 {
    width: 100%;
    height: 100%;
    grid-area: ph4;
    background-color: white;
}
</style>