# SPDX-License-Identifier: GPL-2.0-only
#
# Copyright (C) 2012-2015 OpenWrt.org
# Copyright (C) 2016-2017 LEDE project

include $(TOPDIR)/rules.mk
include $(INCLUDE_DIR)/image.mk

ifeq ($(SUBTARGET),mt7622)
define Image/Build/SDImage
	./make_bundle_image.sh $(KDIR)/mtk-bpi-r64-SD.img \
			       $(STAGING_DIR_IMAGE)/GPT_SD \
			       $(STAGING_DIR_IMAGE)/mtk-bpi-r64-preloader-sd.bin \
			       $(STAGING_DIR_IMAGE)/bpi-r64-atf.img \
			       $(STAGING_DIR_IMAGE)/mtk-bpi-r64-uboot-sd.bin \
			       $(STAGING_DIR_IMAGE)/MT7622_WMAC_EEPROM.bin \
			       $(KDIR)/BPI-R64-SD-kernel.bin \
			       $(KDIR)/root.squashfs
endef

define Image/Build/EMMCImage
	./make_bundle_image.sh $(KDIR)/mtk-bpi-r64-EMMC.img \
			       $(STAGING_DIR_IMAGE)/GPT_EMMC \
			       $(STAGING_DIR_IMAGE)/mtk-bpi-r64-preloader-emmc.bin \
			       $(STAGING_DIR_IMAGE)/bpi-r64-atf.img \
			       $(STAGING_DIR_IMAGE)/mtk-bpi-r64-uboot-emmc.bin \
			       $(STAGING_DIR_IMAGE)/MT7622_WMAC_EEPROM.bin \
			       $(KDIR)/BPI-R64-EMMC-kernel.bin \
			       $(KDIR)/root.squashfs
endef
endif

define Build/sysupgrade-emmc
	rm -f $@.recovery
	mkfs.fat -C $@.recovery 3070

	./gen_$(SUBTARGET)_emmc_img.sh $@ \
		$(IMAGE_KERNEL) \
		$@.recovery \
		$(IMAGE_ROOTFS)
endef

# default all platform image(fit) build 
define Device/Default
  PROFILES = Default $$(DEVICE_NAME)
  KERNEL_NAME := Image
  KERNEL = kernel-bin | lzma | \
	fit lzma $$(KDIR)/image-$$(firstword $$(DEVICE_DTS)).dtb
  KERNEL_INITRAMFS = kernel-bin | lzma | \
	fit lzma $$(KDIR)/image-$$(firstword $$(DEVICE_DTS)).dtb
  FILESYSTEMS := squashfs
  DEVICE_DTS_DIR := $(DTS_DIR)
  IMAGES := sysupgrade.bin
  IMAGE/sysupgrade.bin := append-kernel | pad-to 128k | append-rootfs | \
	pad-rootfs | append-metadata
endef

include $(SUBTARGET).mk

define Image/Build
	$(call Image/Build/$(1),$(1))
	$(call Image/Build/SDImage,$(1))
	$(call Image/Build/EMMCImage,$(1))
endef

$(eval $(call BuildImage))

