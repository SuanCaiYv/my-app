<template>
    <div class="header">
        <PH1></PH1>
        <Home></Home>
        <PH2></PH2>
        <PH3></PH3>
        <Search></Search>
        <PH4></PH4>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Home from "./HomePage.vue"
import Search from "./Search.vue"
import PH1 from "../../placeholder/PH1.vue"
import PH2 from "../../placeholder/PH2.vue"
import PH3 from "../../placeholder/PH3.vue"
import PH4 from "../../placeholder/PH4.vue"

const name = ref<String>("Header")
</script>

<style scoped>
.header {
    width: 100%;
    height: 100%;
    grid-area: header;
    display: grid;
    grid-template-areas: "ph1 home ph2 ph3 search ph4";
    grid-auto-columns: 25px 150px 25px 1fr 200px 60px;
}
</style>