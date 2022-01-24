<template>
    <div class="header">
        <HomePage></HomePage>
        <PH1></PH1>
        <Search></Search>
        <SignButton></SignButton>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import HomePage from "./HomePage.vue"
import Search from "./Search.vue"
import PH1 from "../../placeholder/PH1.vue"
import PH2 from "../../placeholder/PH2.vue"
import PH3 from "../../placeholder/PH3.vue"
import PH4 from "../../placeholder/PH4.vue"
import SignButton from "./SignButton.vue";

const name = ref<String>("Header")
</script>

<style scoped>
.header {
    width: 100%;
    height: 100%;
    grid-area: header;
    display: grid;
    grid-template-areas: "home-page ph1 search sign-button";
    grid-auto-columns: 200px 1fr 200px 200px;
}
</style>