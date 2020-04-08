#!/usr/bin/env bash
echo "Downloading Kubectl Binary..."
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
sudo chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl
echo "Kubectl Client Version..."
kubectl version --client

echo "Downloading Minikube Binary..."
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 \
  && chmod +x minikube

sudo mkdir -p /usr/local/bin/
sudo install minikube /usr/local/bin/

echo "Minikube Version..."

minikube version

sudo minikube start --vm-driver=none --kubernetes-version=v1.16.0