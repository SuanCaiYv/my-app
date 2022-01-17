<template>
    <div class="header">
        <slot>
            <Logo></Logo>
            <Home></Home>
            <PH1></PH1>
            <Search></Search>
            <PH2></PH2>
        </slot>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Logo from "./Logo.vue"
import Home from "./Home.vue"
import PH1 from "./PH1.vue"
import Search from "./Search.vue"
import PH2 from "./PH2.vue"

const name = ref<String>("Header")
</script>

<style scoped>
.header {
    width: 100%;
    height: 100%;
    display: grid;
    grid-template-areas: "logo home ph1 search ph2";
    grid-auto-columns: 120px 120px 1fr 200px 60px;
}
</style>