from __future__ import unicode_literals

from django.db import models


class Ninja(models.Model):
    ninja_id = models.AutoField(primary_key=True)
    first_name = models.TextField()
    last_name = models.TextField()
    sex = models.CharField(max_length=1)
    age = models.IntegerField(blank=True, null=True)
    occupation = models.TextField()
    instagram = models.TextField()
    twitter = models.TextField()

    def __unicode__(self):
        return '{0} {1}'.format(self.first_name, self.last_name)

    class Meta:
        managed = False
        db_table = 'ninja'


class Course(models.Model):
    course_id = models.AutoField(primary_key=True)
    city = models.TextField()
    category = models.TextField()
    season = models.IntegerField()
    size = models.IntegerField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'course'


class Obstacle(models.Model):
    obstacle_id = models.AutoField(primary_key=True)
    title = models.TextField()
    course = models.ForeignKey(
        Course, models.DO_NOTHING, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'obstacle'


class ObstacleResult(models.Model):
    result_id = models.AutoField(primary_key=True)
    transition = models.DecimalField(max_digits=65535, decimal_places=65535)
    duration = models.DecimalField(
        max_digits=65535, decimal_places=65535, blank=True, null=True)
    completed = models.BooleanField()
    obstacle = models.ForeignKey(
        Obstacle, models.DO_NOTHING, blank=True, null=True)
    ninja = models.ForeignKey(
        Ninja, models.DO_NOTHING, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'obstacleresult'


class CourseResult(models.Model):
    result_id = models.AutoField(primary_key=True)
    duration = models.DecimalField(
        max_digits=65535, decimal_places=65535, blank=True, null=True)
    finish_point = models.IntegerField()
    completed = models.BooleanField()
    course = models.ForeignKey(
        Course, models.DO_NOTHING, blank=True, null=True)
    ninja = models.ForeignKey(
        Ninja, models.DO_NOTHING, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'courseresult'


class CareerSummary(models.Model):
    summary_id = models.AutoField(primary_key=True)
    best_finish = models.TextField()
    speed = models.DecimalField(max_digits=65535, decimal_places=65535)
    success = models.DecimalField(max_digits=65535, decimal_places=65535)
    consistency = models.DecimalField(max_digits=65535, decimal_places=65535)
    seasons = models.IntegerField()
    qualifying = models.IntegerField()
    finals = models.IntegerField()
    stages = models.IntegerField()
    ninja = models.ForeignKey(
        Ninja, models.DO_NOTHING, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'careersummary'
