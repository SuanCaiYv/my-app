<template>
    <div class="ph3"></div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("PH3")
</script>

<style scoped>
.ph3 {
    width: 100%;
    height: 100%;
    grid-area: ph3;
    background-color: white;
}
</style>