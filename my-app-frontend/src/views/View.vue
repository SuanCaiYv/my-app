<template>
    <div class="view">
        <Layout></Layout>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Layout from "../components/view/Layout.vue"

const name = ref<string>("View")
</script>

<style scoped>
.view {
    width: 100%;
    height: 100%;
}
</style>