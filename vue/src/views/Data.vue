<template>
	<div class="container">
  		<div class="enter-address-wrapper">
          <div class="input-label">Search Social Data & Transfer</div>

          <div class="input-wrapper">
            <input
              v-model="state.address"
              class="input"
              :class="{
                error: state.address.length > 0 && !validAddress
              }"
              placeholder="Mises ID"
            />
            <div
              v-if="state.address.length > 0 && !validAddress"
              class="error-message"
            >
              Invalid address
            </div>
          </div>
        </div>
	</div>
	<SpCrudReadonly store-name="misesid.misestm.v1beta1" item-name="UserInfo" :address="searchAddress" display-name="User Info" />
	<!-- <SpCrudReadonly store-name="misesid.misestm.v1beta1" item-name="UserRelation" :address="searchAddress"  /> -->
	<SpCrudReadonly store-name="cosmos.tx.v1beta1" item-name="Tx" :address="searchAddress"  display-name="Transfer Recv" />
	<SpCrudReadonly store-name="cosmos.tx.v1beta1" item-name="Tx" :address="searchAddress"  display-name="Transfer Send" />
	<!-- <SpCrudReadonly store-name="misesid.misestm.v1beta1" item-name="AppInfo"  address="searchAddress" />
	<SpCrudReadonly store-name="misesid.misestm.v1beta1" item-name="DidRegistry" address="searchAddress"  /> -->
</template>

<script lang="ts">
import SpCrudReadonly from '../components/SpCrudReadonly'
import { Bech32 } from '@cosmjs/encoding'
import { computed, reactive} from 'vue'
export interface State {
  address: string
}

export let initialState: State = {
  address: '',
}
export default {
	components: {
		SpCrudReadonly
	},
	name: 'Data',
	setup() {
		let state: State = reactive(initialState)
		let validAddress = computed<boolean>(() => {
			let valid: boolean

			try {
				valid = !!Bech32.decode(state.address)
			} catch {
				valid = false
			}

			return valid
		})
		let searchAddress = computed<string>(() => {
			let valid: boolean

			try {
				valid = !!Bech32.decode(state.address)
			} catch {
				valid = false
			}
			if (valid) {
				return state.address
			}
			return ""
		})
		return {
			state,
			validAddress,
			searchAddress
		}
	}
}
</script>

<style scoped lang="scss">
.advanced-label {
  font-family: Inter;
  font-style: normal;
  font-weight: 600;
  font-size: 13px;
  line-height: 153.8%;
  /* identical to box height, or 20px */

  display: inline-flex;
  align-items: center;

  /* base/black 0 */

  color: #000000;

  &--disabled {
    color: rgba(0, 0, 0, 0.33);
    &:hover {
      cursor: default !important;
    }
  }
}

.advanced-label:hover {
  cursor: pointer;
}
.copy {
  padding: 12px 0;
}
.feedback {
  display: flex;
  flex-direction: column;
  align-items: center;
}
.tx {
  padding-bottom: 40px;
}
.token-selector {
  &--main {
    &:deep(.add-token) {
      margin-top: 18px;
    }
  }
}

.advanced {
  &:deep(.add-token) {
    margin-top: 17px;
  }
}

.qrcode-wrapper {
  background: rgba(0, 0, 0, 0.03);
  padding: 36px;
  text-align: center;
}

.address-wrapper {
  padding: 16px;
}

.receive-wrapper .address {
  font-family: Inter;
  font-style: normal;
  font-weight: normal;
  font-size: 16px;
  /* or 24px */

  display: flex;
  align-items: center;

  word-break: break-all;

  color: #000000;
}

.tx-feedback-title {
  font-family: Inter;
  font-style: normal;
  font-weight: bold;
  font-size: 21px;
  line-height: 152%;
  /* identical to box height, or 32px */

  text-align: center;
  letter-spacing: -0.017em;

  /* light/text */

  color: #000000;
}
.tx-feedback-subtitle.amount {
  text-transform: uppercase;
}
.tx-feedback-subtitle {
  font-family: Inter;
  font-style: normal;
  font-weight: normal;
  font-size: 16px;
  line-height: 150%;
  /* identical to box height, or 24px */

  text-align: center;

  /* light/muted */

  color: rgba(0, 0, 0, 0.667);
}

.tx-ongoing-title {
  font-family: Inter;
  font-style: normal;
  font-weight: normal;
  font-size: 16px;
  line-height: 150%;
  /* identical to box height, or 24px */

  text-align: center;

  /* light/muted */

  color: rgba(0, 0, 0, 0.667);
}

.tx-ongoing-subtitle {
  font-family: Inter;
  font-style: normal;
  font-weight: bold;
  font-size: 21px;
  line-height: 152%;
  /* identical to box height, or 32px */

  text-align: center;
  letter-spacing: -0.017em;

  /* light/text */

  color: #000000;
}

.title-wrapper {
  display: flex;
}

.input-label {
  font-family: Inter;
  font-style: normal;
  font-weight: 400;
  font-size: 13px;
  line-height: 153.8%;
  /* identical to box height, or 20px */

  /* light/muted */

  color: rgba(0, 0, 0, 0.667);
}

.title {
  font-family: Inter;
  font-style: normal;
  font-weight: bold;
  font-size: 28px;
  line-height: 127%;
  /* identical to box height, or 36px */

  letter-spacing: -0.016em;
  font-feature-settings: 'zero';

  color: rgba(0, 0, 0, 0.33);

  &.disabled {
    &:hover {
      cursor: text;
    }
  }
}

.title.active {
  color: #000000;
}

.title.active:hover {
  cursor: initial;
}

.title:hover {
  cursor: pointer;
}

.enter-address-wrapper {
  display: flex;
  flex-direction: column;
  flex: 1;
}

.input {
  margin-top: 4px;
  padding: 12px 16px;
  height: 48px;
  background-color: rgba(0, 0, 0, 0.03);
  border-radius: 10px;
  font-family: Inter;
  font-style: normal;
  font-weight: normal;
  font-size: 16px;
  line-height: 130%;
  color: #000000;
  width: 100%;
  outline: 0;
  transition: background-color 0.2s cubic-bezier(0.645, 0.045, 0.355, 1);
  display: block;

  &:not([disabled]) {
    &:hover {
      background: rgba(0, 0, 0, 0.07);
    }
  }

  &:focus {
    background: rgba(0, 0, 0, 0.07);
    color: #000;
  }

  &.error {
    box-shadow: 0 0 0 1px rgba(254, 71, 95, 1);
  }
}

.error-message {
  font-family: Inter;
  font-style: normal;
  font-weight: normal;
  font-size: 13px;
  line-height: 153.8%;
  color: #d80228;
  margin-top: 5px;
}

.input::placeholder {
  color: rgba(0, 0, 0, 0.33);
}

.input-wrapper {
  display: block;
}
</style>