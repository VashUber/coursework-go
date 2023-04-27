<script setup lang="ts">
import "vue3-carousel/dist/carousel.css";
import { Carousel, Slide } from "vue3-carousel";

const carouselSlides = [
  {
    src: "/slide-1.jpg",
    alt: "slide1",
  },
  {
    src: "/slide-2.jpg",
    alt: "slide2",
  },
  {
    src: "/slide-3.jpg",
    alt: "slide3",
  },
];
</script>

<template>
  <div class="-mx-4 page p-0">
    <carousel :wrap-around="true" :autoplay="4000" class="aspect-video relative">
      <slide v-for="slide in carouselSlides" :key="slide.alt">
        <img :src="slide.src" :alt="slide.alt" class="w-full aspect-video" />
      </slide>

      <template #addons>
        <div class="bg-black bg-opacity-60 absolute w-full top-0 h-full z-20">
          <div class="flex h-full flex-col justify-center items-center text-white">
            <h1 class="text-3xl">
              {{ $t("page.index.h1") }}
            </h1>
            <router-link to="/tickets">
              <button class="button button--accent">{{ $t("misc.buy") }} {{ $t("misc.ticket") }}</button>
            </router-link>
          </div>
        </div>
      </template>
    </carousel>
  </div>
</template>
