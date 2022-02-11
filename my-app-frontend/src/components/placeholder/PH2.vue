<template>
    <div class="ph2"></div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("PH2")
</script>

<style scoped>
.ph2 {
    width: 100%;
    height: 100%;
    grid-area: ph2;
    background-color: white;
}
</style>