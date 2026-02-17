#!/bin/bash

# Copies localhost's ~/.kube/config file into the container and swaps out localhost
# for host.docker.internal whenever a new shell starts to keep them in sync.
if [ "$SYNC_LOCALHOST_KUBECONFIG" = "true" ] && [ -d "/usr/local/share/kube-localhost" ]; then
    mkdir -p $HOME/.kube
    
    # Check if there are any files to copy
    if [ -n "$(ls -A /usr/local/share/kube-localhost 2>/dev/null)" ]; then
        sudo cp -r /usr/local/share/kube-localhost/. $HOME/.kube
        sudo chown -R $(id -u) $HOME/.kube
        
        if [ -f "$HOME/.kube/config" ]; then
            # Replace localhost/127.0.0.1 with host.docker.internal for kind/docker-desktop
            sed -i -e "s/localhost/host.docker.internal/g" "$HOME/.kube/config"
            sed -i -e "s/127.0.0.1/host.docker.internal/g" "$HOME/.kube/config"
            
            # Fix TLS verification error: since the cert is valid for 'localhost' but not 'host.docker.internal',
            # we must skip verification. We replace the CA data/path with the skip flag.
            sed -i -e "s/certificate-authority-data:.*$/insecure-skip-tls-verify: true/g" "$HOME/.kube/config"
            sed -i -e "s/certificate-authority:.*$/insecure-skip-tls-verify: true/g" "$HOME/.kube/config"
        fi
    fi
fi
