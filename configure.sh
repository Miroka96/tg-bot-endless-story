#!/bin/sh

configdir="conf/"
apikeyprefix="api-key."
apikeysuffix=".conf"


# gets parameter 'target'
function storeApiKey {
	echo -n "Please enter the Telegram API key for Target '${target}': "
	read key
	targetfile=${configdir}${apikeyprefix}${target}${apikeysuffix}
	echo $key > ${targetfile}
	echo "Wrote api key to ${targetfile}"
	echo
}

# gets parameter 'target'
function requestApiKey {
	echo "A Telegram API key is needed"
	echo -n "Target Name for the new API [${target}]: "
	read newtarget

	if [ "${newtarget}" != "" ] ; then 
		target=${newtarget}
	fi

	targetfile=${configdir}${apikeyprefix}${target}${apikeysuffix}

	if [ -e ${targetfile} ] ; then 
		echo -n "API key file '${targetfile}' already exists. Overwrite? [Yn] "
		read overwrite
		if [ "${overwrite}" != "n" ] ; then
			echo "Overwriting ${targetfile} ..."
		else 
			echo "Continuing..."
			echo
			return
		fi
	fi

	storeApiKey
}

echo "Configuring Telegram Bot"
echo

target="deployment"
requestApiKey

target="testing"
requestApiKey

echo "Finished!"