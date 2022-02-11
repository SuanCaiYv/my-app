<template>
    <div class="ph1"></div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("PH1")
</script>

<style scoped>
.ph1 {
    width: 100%;
    height: 100%;
    grid-area: ph1;
    background-color: white;
}
</style>