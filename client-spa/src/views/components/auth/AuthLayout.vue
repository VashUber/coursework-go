<template>
  <div
    class="w-full flex items-center justify-center gap-40 h-[calc(100vh_-_60px)] mx-4"
  >
    <form class="flex-grow flex flex-col gap-2 items-center" @submit.prevent>
      <slot name="form" />
    </form>
    <div
      class="flex-grow aspect-[2/3] overflow-hidden h-[calc(100vh_-_60px)] relative"
    >
      <h1
        class="absolute text-white z-20 top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2"
      >
        <slot name="title" />
      </h1>
      <img
        src="/auth-page-img.jpeg"
        class="w-full h-full object-cover brightness-[60%]"
      />
    </div>
  </div>
</template>
