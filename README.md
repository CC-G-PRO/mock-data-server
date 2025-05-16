api server 에서 제공하는 api 테스트용 mock data 를 제공하는 서버입니다. 
실제 로직 구현은 없으면 단순히 mock data 를  반환하는 식입니다. 

docker build .  
docker run -d -p 80:8080(포트 자유롭게) <image>
