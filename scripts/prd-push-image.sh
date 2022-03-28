#!/bin/bash
function gacha-login-aws() {
  aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin   201591287706.dkr.ecr.ap-northeast-1.amazonaws.com
}
COMMIT=`git rev-parse HEAD`
docker-compose build prd && \
gacha-login-aws && \
  docker tag gacha_prd:latest 201591287706.dkr.ecr.ap-northeast-1.amazonaws.com/gacha/prd:${COMMIT} && \
  docker tag gacha_prd:latest 201591287706.dkr.ecr.ap-northeast-1.amazonaws.com/gacha/prd:latest && \
  docker push 201591287706.dkr.ecr.ap-northeast-1.amazonaws.com/gacha/prd:${COMMIT} && \
  docker push 201591287706.dkr.ecr.ap-northeast-1.amazonaws.com/gacha/prd:latest