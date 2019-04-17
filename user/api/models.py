from django.db import models
from django.contrib.auth.models import (
    BaseUserManager,
    AbstractUser)

# Create your models here.


class ClientManager(BaseUserManager):
    def create_user(self, email, password=None):
        """
        Creates and saves User with given email and password
        :param email:
        :param password:
        :return:
        """

        if not email:
            raise ValueError("No email")

        user = self.model(
            email=self.normalize_email(email),
        )
        user.username = email
        user.set_password(password)
        user.is_active = False
        user.save(using=self._db)
        return user

    def create_superuser(self, email, password, location):
        user = self.create_user(
            email,
            password = password
        )  
        user.location = location
        user.is_admin = True
        user.save(using=self._db)
        return user

class Client(AbstractUser):

    email = models.EmailField(
        verbose_name='email adress',
        max_length= 255,
        unique=True,
    )
    location = models.CharField(max_length=500, default="")

    is_active = models.BooleanField(default=True)
    is_admin = models.BooleanField(default=False)
    objects = ClientManager()

    USERNAME_FIELD = 'email'
    REQUIRED_FIELDS = ['location']

    def __str__(self):
        return self.email

    def has_perm(self, perm, obj=None):
        "Does the user have a specific permission?"
        # Simplest possible answer: Yes, always
        return True

    def has_module_perms(self, app_label):
        "Does the user have permissions to view the app `app_label`?"
        # Simplest possible answer: Yes, always
        return True

    @property
    def is_staff(self):
        "Is the user a member of staff?"
        # Simplest possible answer: All admins are staff
        return self.is_admin

