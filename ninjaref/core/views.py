from django.shortcuts import render

from .models import CareerSummary


def index(request):
    return render(request, 'core/index.html', {})


def leaderboard(request):
    r = {'rating': 'speed + consistency + success'}
    leaders = CareerSummary.objects.all().extra(select=r).order_by('-rating')
    return render(request, 'core/leaderboard.html', {
        'leaders': leaders[:10],
        'men': leaders.filter(ninja__sex='M')[:10],
        'women': leaders.filter(ninja__sex='F')[:10]
    })
