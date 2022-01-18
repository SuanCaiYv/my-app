<template>
    <div class="home">
        <Layout></Layout>
    </div>
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
.home {
    width: 100%;
    height: 100%;
}
</style>