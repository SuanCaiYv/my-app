<template>
    <div class="header">
        <HomePage></HomePage>
        <PH1></PH1>
        <Backup></Backup>
        <SignButton></SignButton>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import HomePage from "../../components/common/HomePage.vue"
import PH1 from "../../components/placeholder/PH1.vue"
import Backup from "./Backup.vue"
import SignButton from "../../components/common/SignButton.vue"

const name = ref<string>("HeaderWithoutSearch")
</script>

<style scoped>
.header {
    width: 100%;
    height: 60px;
    display: grid;
    grid-template-areas: "home-page ph1 backup sign-button";
    grid-auto-columns: 200px 1fr 200px 200px;
    position: fixed;
    background-color: white;
}
</style>