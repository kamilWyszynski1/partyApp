from django.urls import path, include
from rest_framework import routers
from . import views

router = routers.SimpleRouter()
router.register('clients', views.ClientViewSet)

urlpatterns = [
    path('', include(router.urls)),
    path('check/', views.check_user, name='check'),
    path('activate/', views.activate_user, name='activate')
]