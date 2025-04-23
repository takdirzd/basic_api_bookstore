from django.contrib import admin
from django.urls import path
from . import views  # import views dari folder yang sama

urlpatterns = [
    path('admin/', admin.site.urls),
    path('', views.index, name='index'),  # tambahkan ini
]
