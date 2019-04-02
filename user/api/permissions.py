from rest_framework.permissions import BasePermission


class IsAdminOrIsSelf(BasePermission):
    def has_object_permission(self, request, view, obj):
        try:
            if obj.email == request.user.email:
                return True
        except AttributeError:
            return False
