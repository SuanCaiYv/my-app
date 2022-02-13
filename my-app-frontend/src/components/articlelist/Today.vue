<template>
    <div></div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("Today")
</script>

<style scoped>
.today {
    width: 100%;
    height: 100%;
    grid-area: today;
}
</style>