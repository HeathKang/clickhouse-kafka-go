// add /etc/hosts 127.0.0.1 kuka-connect-kafka-0.kuka-connect-kafka.kuka-connect-dev.svc.digital-dev.kukaplus.com 
kubectl port-forward pods/kuka-connect-kafka-0 9092:9092 -n kuka-connect-dev
