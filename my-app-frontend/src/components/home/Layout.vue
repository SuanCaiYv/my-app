<template>
    <div class="layout">
        <Header></Header>
        <PH1></PH1>
        <Body></Body>
        <Footer></Footer>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Header from "../common/HeaderWithSearch.vue"
import PH1 from "../placeholder/PH1.vue"
import Body from "./Body.vue"
import Footer from "./Footer.vue"

const name = ref<String>("Layout")
</script>

<style scoped>
.layout {
    width: 100%;
    height: 100%;
    display: grid;
    grid-template-areas:
        "ph1"
        "body"
        "footer";
    grid-template-rows: 60px 1fr 50px;
}
</style>