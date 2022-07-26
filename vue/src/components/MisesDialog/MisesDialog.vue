<template>
  <transition name="dialog-fade" @after-enter="afterEnter" @after-leave="afterLeave">
    <div v-show="visible" class="el-dialog__wrapper" @click.self="handleWrapperClick">
      <div role="dialog" :key="key" aria-modal="true" :aria-label="title || 'dialog'" class="el-dialog" ref="dialog" :style="style">
        <div class="el-dialog__header">
          <slot name="title">
            <span class="el-dialog__title">{{ title }}</span>
          </slot>
          <button type="button" class="el-dialog__headerbtn" aria-label="Close" v-if="showClose" @click="handleClose">
            <svg
              t="1650446111880"
              class="icon el-dialog__close"
              viewBox="0 0 1024 1024"
              version="1.1"
              xmlns="http://www.w3.org/2000/svg"
              p-id="4054"
              width="20"
              height="20"
            >
              <path
                d="M463.787 512l-253.44-253.44a34.133 34.133 0 0 1 48.213-48.213L512 463.787l253.44-253.44a34.133 34.133 0 1 1 48.213 48.213L560.213 512l253.44 253.44a34.133 34.133 0 1 1-48.213 48.213L512 560.213l-253.44 253.44a34.133 34.133 0 0 1-48.213-48.213z"
                p-id="4055"
              ></path>
            </svg>
            <!-- <i class="el-dialog__close el-icon el-icon-close"></i> -->
          </button>
        </div>
        <div class="el-dialog__body"><slot></slot></div>
        <div class="el-dialog__footer" v-if="$slots.footer">
          <slot name="footer"></slot>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
import { computed, defineComponent } from 'vue'
export default defineComponent({
  name: 'MisesDialog',
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    showClose: {
      type: Boolean,
      default: true
    },
    title: {
      type: String,
      default: ''
    },
    top: {
      type: String,
      default: '15vh'
    },
    width: String
  },
  setup(props, { emit }) {
    const handleWrapperClick = (e) => {
      if (e.target !== e.currentTarget) return
      handleClose()
    }
    const handleClose = () => emit('update:visible', false)

    const afterEnter = () => emit('opened')
    const afterLeave = () => emit('closed')
    const style = computed(() => {
      let style = {}
      style.marginTop = props.top
      if (props.width) {
        style.width = props.width
      }
      return style
    })
    return {
      handleWrapperClick,
      handleClose,
      afterEnter,
      afterLeave,
      style
    }
  }
})
</script>
<style lang="scss" scoped>
.el-dialog {
  position: relative;
  margin: 0 auto 50px;
  background: #fff;
  border-radius: 2px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.3);
  box-sizing: border-box;
  width: 50%;
}
@media (max-width: 768px) {
	.el-dialog {
		width: 95%;
	}
}
.el-dialog.is-fullscreen {
  width: 100%;
  margin-top: 0;
  margin-bottom: 0;
  height: 100%;
  overflow: auto;
}

.el-dialog__wrapper {
  position: fixed;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  overflow: auto;
  margin: 0;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 66;
}

.el-dialog__header {
  padding: 20px 20px 10px;
}

.el-dialog__headerbtn {
  position: absolute;
  top: 20px;
  right: 20px;
  padding: 0;
  background: transparent;
  border: none;
  outline: none;
  cursor: pointer;
  font-size: 16px;
}

.el-dialog__headerbtn .el-dialog__close {
  color: #909399;
}

.el-dialog__headerbtn:focus .el-dialog__close,
.el-dialog__headerbtn:hover .el-dialog__close {
  color: #409eff;
}

.el-dialog__title {
  line-height: 24px;
  font-size: 18px;
  color: #303133;
}

.el-dialog__body {
  padding: 30px 20px;
  color: #606266;
  font-size: 14px;
  word-break: break-all;
}

.el-dialog__footer {
  padding: 10px 20px 20px;
  text-align: right;
  box-sizing: border-box;
}

.el-dialog--center {
  text-align: center;
}

.el-dialog--center .el-dialog__body {
  text-align: initial;
  padding: 25px 25px 30px;
}

.el-dialog--center .el-dialog__footer {
  text-align: inherit;
}

.dialog-fade-enter-active {
  animation: dialog-fade-in 0.3s;
}

.dialog-fade-leave-active {
  animation: dialog-fade-out 0.3s;
}

@keyframes dialog-fade-in {
  0% {
    transform: translate3d(0, -20px, 0);
    opacity: 0;
  }

  to {
    transform: translateZ(0);
    opacity: 1;
  }
}

@keyframes dialog-fade-out {
  0% {
    transform: translateZ(0);
    opacity: 1;
  }

  to {
    transform: translate3d(0, -20px, 0);
    opacity: 0;
  }
}
</style>