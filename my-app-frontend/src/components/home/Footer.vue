<template>
    <div class="footer">©️ CodeWithBuff</div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("Footer")
</script>

<style scoped>
.footer {
    width: 100%;
    height: 100%;
    grid-area: footer;
}
</style>