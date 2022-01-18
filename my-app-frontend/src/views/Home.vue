<template>
    <Layout class="layout">
    </Layout>
</template>

<script setup lang="ts">
import {provide, ref} from "vue"
import Layout from "../components/home/layout/Layout.vue"

const name = ref<String>("Home")

provide("searchKey", function (key: String) {
    console.log(key)
})
</script>

<style scoped>
.layout {
    width: 100%;
    height: 100%;
}
</style>