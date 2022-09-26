#KERNEL_LOADADDR := 0x44080000
KERNEL_LOADADDR := 0x48080000    ## bpi-r64

define Device/BPI-R64-SD
  DEVICE_VENDOR := Banana Pi
  DEVICE_MODEL := Banana Pi R64
  DEVICE_TITLE := MTK7622 BPI R64 SD
  DEVICE_DTS := mt7622-bananapi-bpi-r64-sd
  DEVICE_DTS_DIR := $(DTS_DIR)/mediatek
  DEVICE_PACKAGES := kmod-usb-ohci kmod-usb2 kmod-usb3 \
  		     kmod-ata-core kmod-ata-ahci-mtk \
		     kmod-mt7615e mt7622bt-firmware swconfig
  SUPPORTED_DEVICES := bananapi,bpi-r64
endef
TARGET_DEVICES += BPI-R64-SD

define Device/BPI-R64-EMMC
  DEVICE_VENDOR := Banana Pi
  DEVICE_MODEL := Banana Pi R64
  DEVICE_TITLE := MTK7622 BPI R64 EMMC
  DEVICE_DTS := mt7622-bananapi-bpi-r64-emmc
  DEVICE_DTS_DIR := $(DTS_DIR)/mediatek
  DEVICE_PACKAGES := kmod-usb-ohci kmod-usb2 kmod-usb3 \
  		     kmod-ata-core kmod-ata-ahci-mtk \
                     kmod-mt7615e mt7622bt-firmware swconfig
  SUPPORTED_DEVICES := bananapi,bpi-r64
endef
TARGET_DEVICES += BPI-R64-EMMC


