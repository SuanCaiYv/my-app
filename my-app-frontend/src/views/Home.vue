<template>
    <Layout class="layout">
        <Header></Header>
        <Body></Body>
        <Footer></Footer>
    </Layout>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Layout from "../components/home/Layout.vue"
import Header from "../components/home/Header.vue"
import Body from "../components/home/Body.vue"
import Footer from "../components/home/Footer.vue"

const name = ref<String>("Home")
</script>

<style scoped>
.layout {
    width: 100%;
    height: 100%;
}
</style>