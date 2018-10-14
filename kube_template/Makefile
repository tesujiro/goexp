apply_deployment:
	kubectl apply -f deployment.yaml

apply_service:
	kubectl apply -f service.yaml

make_index:
	for PODNAME in `kubectl get pods -l app=sample-app -o jsonpath='{.items[*].metadata.name}'`; do \
		kubectl exec -it $${PODNAME} -- cp /etc/hostname /usr/share/nginx/html/index.html ;\
	done

access:
	kubectl run --image=centos:6 --restart=Never --rm -i testpod -- curl -s http://`kubectl get svc|grep sample-clusterip|awk '{print $$3}'`:8080

drop_all:
	kubectl delete svc sample-clusterip
	kubectl delete deployment sample-deployment 
